import {
    ContainerColumn,
    ContainerColumnCenter,
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {useState} from "react";
import ProfileButtonsContainer
    from "@/components/UI/ProfilePage/Profile/ProfileButtonsContainer/ProfileButtonsContainer";
import ProfileContainer from "@/components/UI/ProfilePage/Profile/ProfileContainer/ProfileContainer";

export default function ProfileComponent() {
    const [avatar, setAvatar] = useState("/images/dada2.jpg")
    const [name, setName] = useState("Григорий")
    const [surname, setSurname] = useState("Перфилин")
    const [email, setEmail] = useState("grigorij.perfilin@mail.ru")
    const [description, SetDescription] = useState("«Я уже говорил тебе, что такое безумие ? Безумие — это точное повторение одного и того же действия, раз за разом, в надежде на изменение»...\n" +
        "\n" +
        "О, пудж не в бане!!!")

    return (
        <ContainerColumnCenter>
            <ContainerColumn style={{width: "calc(1200rem/16)", height: "calc(624rem/16)"}}>
                <ProfileButtonsContainer/>
                <div style={{width: "100%", height: "calc(100% - 96rem / 16 - 6rem/16)"}}>
                    <ProfileContainer avatarSrc={avatar} name={name} surname={surname} email={email} description={description}/>
                </div>
            </ContainerColumn>
        </ContainerColumnCenter>
    )
}