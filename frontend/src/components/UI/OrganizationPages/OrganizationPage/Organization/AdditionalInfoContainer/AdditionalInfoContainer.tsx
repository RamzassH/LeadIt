import {
    Background,
    BackgroundContainer,
    ContainerRow
} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Title";
import {Text} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Text";
import React, {useRef, useState} from "react";
import useOrganizationInfoStore from "@/components/UI/OrganizationPages/OrganizationPage/store";
import EditDescriptionDataWindow
    from "@/components/UI/OrganizationPages/OrganizationPage/ModalWindow/EditDescriptionDataWindow/EditDescriptionDataWindow";

interface AdditionalInfoProps {
    style?: React.CSSProperties;
}

export default function AdditionalInfoContainer({style}: AdditionalInfoProps) {
    const [open, setOpen] = useState(false);
    const description = useOrganizationInfoStore(state => state.info.description);

    const handleOpen = () => {setOpen(true);};
    const handleClose = () => {setOpen(false);};

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
                                Описание
                            </Title>
                            <div style={{marginRight: "auto"}}/>
                            <ChangeButton onClick={handleOpen}>
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
            <EditDescriptionDataWindow open={open} handleClose={handleClose}/>
        </div>
    )
}