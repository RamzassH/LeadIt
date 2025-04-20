import {
    ContainerColumnCenter
} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Containers";
import {useEffect, useState} from "react";
import OrganizationContainer from "@/components/UI/OrganizationPages/OrganizationPage/Organization/OrganizationContainer/OrganizationContainer";
import useOrganizationInfoStore from "@/components/UI/OrganizationPages/OrganizationPage/store";

export default function OrganizationComponent() {
    const setLogo = useOrganizationInfoStore(state => state.setLogo);
    const setOrganization = useOrganizationInfoStore(state => state.setOrganization);
    const setContactInfo = useOrganizationInfoStore(state => state.setContactInfo);
    const setDescription = useOrganizationInfoStore(state => state.setDescription);

    useEffect(() => {
        setLogo({
            src: "/images/dada2.jpg",
            positionX: 0,
            positionY: 0,
        })
        setOrganization({
            name: "ООО \"Контора пидорасов\"",
            director: "Залупов Залуп Залупович",
        })
        setContactInfo({
            email: "grigorij.perfilin@mail.ru",
            messenger: "tg: @yanaCist",
            phone: "8 800 555 35 35"
        })
        setDescription("Самая лучшая контора на свете. Никаких долбаебов. Только настоящие любители жесткого порева(Игры в доту)")
    }, []);

    return (
        <ContainerColumnCenter>
            <OrganizationContainer
                style={{
                    width: "calc(1400rem/16)", height: "calc(900rem/16)"
                }}
            />
        </ContainerColumnCenter>
    )
}