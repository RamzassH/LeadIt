import {
    Background,
    BackgroundContainer,
    ContainerColumn,
    ContainerRow
} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Title";
import MainInfoContainerComponent
    from "@/components/UI/OrganizationPages/OrganizationPage/Organization/MainInfoContainerComponent/MainInfoContainerComponent";
import useOrganizationInfoStore from "@/components/UI/OrganizationPages/OrganizationPage/store";
import EditOrganizationDataWindow
    from "@/components/UI/OrganizationPages/OrganizationPage/ModalWindow/EditOrganizationDataWindow/EditOrganizationDataWindow";
import {useState} from "react";

interface MainInfoProps {

}

interface ComponentProps {
    name: string;
    value: string;
}

export default function MainInfoContainer({}:MainInfoProps) {
    const info = useOrganizationInfoStore(state => state.info);
    const [open, setOpen] = useState(false);

    const handleOpen = () => {setOpen(true);};
    const handleClose = () => {setOpen(false);};

    const contactInfo: ComponentProps[] = [
        {name: "Наименование", value: info.organization.name},
        {name: "Директор", value: info.organization.director},
        {name: "Email", value: info.contacts.email},
        {name: "Мессенджер", value: info.contacts.messenger},
        {name: "Телефон", value: info.contacts.phone},
    ]

    return (
        <div style={{width: "100%", height: "100%", marginBottom: "calc(6rem/16)"}}>
            <BackgroundContainer style={{padding: "calc(36rem/16)"}}>
                <ContainerColumn style={{gap: "calc(36rem/16)"}}>
                    <Background>
                        <ContainerColumn>
                            <div style={{padding: "calc(12rem/16) calc(16rem/16)"}}>
                                <ContainerRow>
                                    <Title>
                                        Контактная информация
                                    </Title>
                                    <div style={{marginRight: "auto"}}/>
                                    <ChangeButton onClick={handleOpen}>
                                        изменить
                                    </ChangeButton>
                                </ContainerRow>
                            </div>
                            <ContainerColumn style={{padding: "calc(0rem/16) calc(16rem/16) calc(12rem/16) calc(16rem/16)", gap: ".5rem"}}>
                                {contactInfo.map((item, index) => (
                                    <MainInfoContainerComponent label={item.name} text={item.value} key={index}/>
                                ))}
                            </ContainerColumn>
                        </ContainerColumn>
                    </Background>
                </ContainerColumn>
            </BackgroundContainer>
            <EditOrganizationDataWindow open={open} handleClose={handleClose}/>
        </div>
    )
}