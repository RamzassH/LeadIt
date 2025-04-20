import {Box, Paper, Typography } from "@mui/material";
import { Column } from "../ScrumBoard/styled/styled";
import TextField from "@mui/material/TextField";
import {useState} from "react";

const RetrospectiveBoard: React.FC = () => {
    const [feedback, setFeedback] = useState<{
        good: string[];
        bad: string[];
        actions: string[];
    }>({ good: [], bad: [], actions: [] });
    const [newFeedback, setNewFeedback] = useState('');

    const addFeedback = (type: keyof typeof feedback) => {
        if (newFeedback.trim()) {
            setFeedback({
                ...feedback,
                [type]: [...feedback[type], newFeedback],
            });
            setNewFeedback('');
        }
    };

    return (
        <Paper sx={{ p: 3, mt: 3 }}>
            <Typography variant="h6" gutterBottom>
                Sprint Retrospective
            </Typography>
            <Box display="flex" justifyContent="space-between" mt={2}>
                <Column sx={{ width: '30%' }}>
                    <Typography color="success.main">What went well</Typography>
                    {feedback.good.map((item, i) => (
                        <Typography key={i}>{item}</Typography>
                    ))}
                    <TextField
                        size="small"
                        placeholder="Add positive feedback"
                        value={newFeedback}
                        onChange={(e) => setNewFeedback(e.target.value)}
                        onKeyPress={(e) => e.key === 'Enter' && addFeedback('good')}
                    />
                </Column>
                <Column sx={{ width: '30%' }}>
                    <Typography color="error.main">What didn't go well</Typography>
                    {feedback.bad.map((item, i) => (
                        <Typography key={i}>{item}</Typography>
                    ))}
                    <TextField
                        size="small"
                        placeholder="Add negative feedback"
                        value={newFeedback}
                        onChange={(e) => setNewFeedback(e.target.value)}
                        onKeyPress={(e) => e.key === 'Enter' && addFeedback('bad')}
                    />
                </Column>
                <Column sx={{ width: '30%' }}>
                    <Typography color="info.main">Action items</Typography>
                    {feedback.actions.map((item, i) => (
                        <Typography key={i}>{item}</Typography>
                    ))}
                    <TextField
                        size="small"
                        placeholder="Add action item"
                        value={newFeedback}
                        onChange={(e) => setNewFeedback(e.target.value)}
                        onKeyPress={(e) => e.key === 'Enter' && addFeedback('actions')}
                    />
                </Column>
            </Box>
        </Paper>
    );
};

export default RetrospectiveBoard;