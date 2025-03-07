import {
    Background,
    BackgroundContainer,
    ContainerRow
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/ProfilePage/Profile/styled/Title";
import {Text} from "@/components/UI/ProfilePage/Profile/styled/Text";
import React, {useRef} from "react";
import useUserInfoStore from "@/components/UI/ProfilePage/store";
import EditDescriptionDataWindow
    from "@/components/UI/ProfilePage/ModalWindow/EditDescriptionDataWindow/EditDescriptionDataWindow";
import useDescriptionStore from "@/components/UI/ProfilePage/ModalWindow/EditDescriptionDataWindow/store";

interface AdditionalInfoProps {
    style?: React.CSSProperties;
}

export default function AdditionalInfoContainer({style}: AdditionalInfoProps) {
    const description = useUserInfoStore(state => state.info.description);
    const handleOpen = useDescriptionStore(state => state.handleOpen)
    const formattedText = description.split("\n").map((line, index) => (
        <span key={index}>
            {line}
            <br/>
        </span>
    ));

    return (
        <div style={style}>
            <BackgroundContainer style={{justifyContent: "center", alignItems: "center"}}>
                <Background style={{
                    width: "calc(100% - 36rem/16 * 2)",
                    height: "calc(100% - 36rem/16 * 2)",
                    }}
                >
                    <div style={{padding: "calc(12rem/16) calc(16rem/16)"}}>
                        <ContainerRow>
                            <Title>
                                О себе
                            </Title>
                            <div style={{marginRight: "auto"}}/>
                            <ChangeButton onClick={() => {handleOpen({description: description})}}>
                                изменить
                            </ChangeButton>
                        </ContainerRow>
                    </div>
                        <BackgroundContainer style={{
                            width: "calc(100% - 16rem/16 * 2)",
                            height: "fit-content",
                            margin: "0 calc(16rem/16) calc(16rem/16)",
                            padding: "calc(16rem/16)"
                            }}
                        >
                            <Text>
                                {formattedText}
                            </Text>
                        </BackgroundContainer>
                </Background>
            </BackgroundContainer>
            <EditDescriptionDataWindow/>
        </div>
    )
}