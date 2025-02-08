import {
    BackgroundContainer,
    ContainerColumn,
    ContainerRow
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {ChangeButton, Title} from "@/components/UI/ProfilePage/Profile/styled/Title";
import {Text} from "@/components/UI/ProfilePage/Profile/styled/Text";
import AvatarProfileContainer from "@/components/UI/ProfilePage/Profile/AvatarProfileContainer/AvatarProfileContainer";
import MainInfoContainer from "@/components/UI/ProfilePage/Profile/MainInfoContainer/MainInfoContainer";
import AdditionalInfoContainer
    from "@/components/UI/ProfilePage/Profile/AdditionalInfoContainer/AdditionalInfoContainer";

interface ProfileContainerProps {
    avatarSrc: string;
    name: string;
    surname: string;
    email: string;
    description: string;
}

export default function ProfileContainer({avatarSrc, name, surname, email, description}:ProfileContainerProps) {

    return (
        <ContainerRow>
            <AvatarProfileContainer src={avatarSrc}/>
            <div style={{width: "calc(100% - 6rem/16 - 435%/1200 * 100)", height: "100%"}}>
                <ContainerColumn>
                    <MainInfoContainer name={name} surname={surname} email={email}/>
                    <AdditionalInfoContainer description={description}/>
                </ContainerColumn>
            </div>
        </ContainerRow>
    )
}