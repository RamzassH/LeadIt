import {
    ContainerColumnCenter
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {useEffect, useState} from "react";
import ProfileContainer from "@/components/UI/ProfilePage/Profile/ProfileContainer/ProfileContainer";
import useUserInfoStore from "@/components/UI/ProfilePage/store";

export default function ProfileComponent() {
    const setAvatar = useUserInfoStore(state => state.setAvatar);
    const setFullName = useUserInfoStore(state => state.setFullName);
    const setDate = useUserInfoStore(state => state.setDate);
    const setContactInfo = useUserInfoStore(state => state.setContactInfo);
    const setProjectInfo = useUserInfoStore(state => state.setProjectInfo);
    const setDescription = useUserInfoStore(state => state.setDescription);

    useEffect(() => {
        setAvatar({
            src: "/images/dada2.jpg",
            positionX: 0,
            positionY: 0,
        })
        setFullName({
            name: "Григорий",
            surname: "Перфилин",
            patronymic: "Александрович"
        })
        setDate("11.12.2003")
        setContactInfo({
            email: "grigorij.perfilin@mail.ru",
            messenger: "tg: @yanaCist"
        })
        setProjectInfo({
            organization: "БГТУ Шушухова",
            projects: "Лукиту, Технологии Надежности",
            position: "Программист-дебил"
        })
        setDescription("«Я уже говорил тебе, что такое безумие ? Безумие — это точное повторение одного и того же действия, раз за разом, в надежде на изменение»...\n" +
            "\n" +
            "О, пудж не в бане!!!")
    }, []);

    return (
        <ContainerColumnCenter>
            <ProfileContainer
                style={{
                    width: "calc(1400rem/16)", height: "calc(900rem/16)"
                }}
            />
        </ContainerColumnCenter>
    )
}