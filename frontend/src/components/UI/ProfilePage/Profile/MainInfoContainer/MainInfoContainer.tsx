import {
    Background,
    BackgroundContainer,
    ContainerColumn,
    ContainerRow
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/ProfilePage/Profile/styled/Title";
import MainInfoContainerComponent
    from "@/components/UI/ProfilePage/Profile/MainInfoContainerComponent/MainInfoContainerComponent";
import useUserInfoStore from "@/components/UI/ProfilePage/store";

interface MainInfoProps {

}

interface ComponentProps {
    name: string;
    value: string;
}

export default function MainInfoContainer({}:MainInfoProps) {
    const info = useUserInfoStore(state => state.info);

    const contactInfo: ComponentProps[] = [
        {name: "Имя", value: info.fullName.name},
        {name: "Фамилия", value: info.fullName.surname},
        {name: "Отчество", value: info.fullName.patronymic},
        {name: "Дата рождения", value: info.date},
        {name: "Email", value: info.contacts.email},
        {name: "Мессенджер", value: info.contacts.messenger},
    ]
    const projectInfo: ComponentProps[] = [
        {name: "Организация", value: info.projectInfo.organization},
        {name: "Проекты", value: info.projectInfo.projects},
        {name: "Должность", value: info.projectInfo.position}
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
                                    <ChangeButton>
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
                    <Background>
                        <ContainerColumn>
                            <div style={{padding: "calc(12rem/16) calc(16rem/16)"}}>
                                <ContainerRow>
                                    <Title>
                                        Проектная деятельность
                                    </Title>
                                    <div style={{marginRight: "auto"}}/>
                                    <ChangeButton>
                                        изменить
                                    </ChangeButton>
                                </ContainerRow>
                            </div>
                            <ContainerColumn style={{padding: "calc(0rem/16) calc(16rem/16) calc(12rem/16) calc(16rem/16)", gap:".5rem"}}>
                                {projectInfo.map((item, index) => (
                                    <MainInfoContainerComponent label={item.name} text={item.value} key={index}/>
                                ))}
                            </ContainerColumn>
                        </ContainerColumn>
                    </Background>
                </ContainerColumn>
            </BackgroundContainer>
        </div>
    )
}