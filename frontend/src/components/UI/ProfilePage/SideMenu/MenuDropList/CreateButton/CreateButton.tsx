import {DropListBackground} from "@/components/UI/ProfilePage/SideMenu/MenuDropList/styles/DropList";
import {GradientButton} from "@/components/UI/ProfilePage/SideMenu/MenuDropList/CreateButton/styled/CreateButton";

export default function CreateButton(props: {}) {
    return (
        <DropListBackground>
            <GradientButton>
                Создать
            </GradientButton>
        </DropListBackground>
    )
}