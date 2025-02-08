import {
    BackgroundContainer,
    ContainerColumn,
    ContainerRow
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/ProfilePage/Profile/styled/Title";
import {Text} from "@/components/UI/ProfilePage/Profile/styled/Text";

interface MainInfoProps {
    name: string;
    surname: string;
    email: string;
}

export default function MainInfoContainer({name, surname, email}:MainInfoProps) {
    return (
        <div style={{width: "100%", height: "calc(320rem/16)", marginBottom: "calc(6rem/16)"}}>
            <BackgroundContainer>
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
                <ContainerColumn style={{padding: "calc(0rem/16) calc(16rem/16) calc(12rem/16) calc(16rem/16)"}}>
                    <Text style={{width: "100%"}}>
                        Имя
                    </Text>
                    <Text style={{width: "100%", fontWeight: "400", marginBottom: "calc(4rem/16)"}}>
                        {name}
                    </Text>
                    <Text style={{width: "100%"}}>
                        Фамилия
                    </Text>
                    <Text style={{width: "100%", fontWeight: "400", marginBottom: "calc(4rem/16)"}}>
                        {surname}
                    </Text>
                    <Text style={{width: "100%"}}>
                        Email
                    </Text>
                    <Text style={{width: "100%", fontWeight: "400", marginBottom: "calc(4rem/16)"}}>
                        {email}
                    </Text>
                </ContainerColumn>
            </BackgroundContainer>
        </div>
    )
}