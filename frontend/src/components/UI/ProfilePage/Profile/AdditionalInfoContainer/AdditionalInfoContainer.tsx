import {
    BackgroundContainer,
    ContainerColumn,
    ContainerRow
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/ProfilePage/Profile/styled/Title";
import {Text} from "@/components/UI/ProfilePage/Profile/styled/Text";

interface AdditionalInfoProps {
    description: string;
}

export default function AdditionalInfoContainer({description}: AdditionalInfoProps) {
    const formattedText = description.split("\n").map((line, index) => (
        <span key={index}>
            {line}
            <br />
        </span>
    ));

    return (
        <div style={{width: "100%", height: "calc(100% - 320rem/16 - 6rem/16)"}}>
            <BackgroundContainer>
                <div style={{padding: "calc(12rem/16) calc(16rem/16)"}}>
                    <ContainerRow>
                        <Title>
                            Дополнительная информация
                        </Title>
                        <div style={{marginRight: "auto"}}/>
                        <ChangeButton>
                            изменить
                        </ChangeButton>
                    </ContainerRow>
                </div>
                <ContainerColumn style={{padding: "calc(0rem/16) calc(16rem/16) calc(12rem/16) calc(16rem/16)"}}>
                    <Text style={{width: "100%", fontWeight: "400"}}>
                        {formattedText}
                    </Text>
                </ContainerColumn>
            </BackgroundContainer>
        </div>
    )
}