import {
    BackgroundContainer, ContainerColumn,
    ContainerColumnCenter,
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {AvatarContainer} from "@/components/UI/ProfilePage/Profile/styled/Avatar";
import {AvatarEditButton} from "@/components/UI/ProfilePage/Profile/ProfileButton/styled/ProfileButton";

interface AvatarContainerProps {
    src: string;
}

export default function AvatarProfileContainer({src}:AvatarContainerProps) {
    return (
        <div style={{width: "calc(435%/1200 * 100)", height: "100%", marginRight: "calc(6rem / 16)"}}>
            <BackgroundContainer>
                <ContainerColumn>
                    <AvatarContainer>
                        <img src={src} alt="Пахнешь слабостью"/>
                    </AvatarContainer>
                    <div style={{width: "100%", height: "fit-content", display: "flex", justifyContent: "center"}}>
                        <AvatarEditButton onClick={(event) => {}}>
                            Изменить аватар
                        </AvatarEditButton>
                    </div>
                </ContainerColumn>
            </BackgroundContainer>
        </div>
    )
}