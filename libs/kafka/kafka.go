package kafka

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"time"
)

type Message = kafka.Message

type Producer struct {
	writer *kafka.Writer
	logger zerolog.Logger
}

func NewProducer(brokers []string, topic string, logger zerolog.Logger) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			WriteTimeout: 10 * time.Second,
		},
		logger: logger,
	}
}

func (p *Producer) Send(ctx context.Context, key, value []byte) error {
	msg := &kafka.Message{
		Key:   key,
		Value: value,
	}
	err := p.writer.WriteMessages(ctx, *msg)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to write message")

		return err
	}

	p.logger.Info().Msg("KafkaProducer: message sent successfully")

	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

type Consumer struct {
	reader *kafka.Reader
	logger zerolog.Logger
}

func NewConsumer(brokers []string, topic, groupID string, logger zerolog.Logger) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        brokers,
			Topic:          topic,
			GroupID:        groupID,
			MinBytes:       10e3,
			MaxBytes:       10e6,
			CommitInterval: time.Second,
		}),
		logger: logger,
	}
}

func (c *Consumer) ReadMessage(ctx context.Context, handler func(msg kafka.Message) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				c.logger.Info().Msg("Context canceled, stopping message consumption")
				return nil
			}
			c.logger.Error().Err(err).Msg("Failed to read message")
			return err
		}

		c.logger.Info().
			Str("topic", msg.Topic).
			Int("partition", msg.Partition).
			Int64("offset", msg.Offset).
			Msgf("Received message: %s", string(msg.Value))

		if err := handler(msg); err != nil {
			c.logger.Error().Err(err).Msg("Failed to handle message")
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
