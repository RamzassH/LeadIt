import {BackgroundContainer} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Containers";
import {Text} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Text";

interface Props {
    label: string;
    text: string;
}

export default function MainInfoContainerComponent({label, text}: Props) {
    return (
        <BackgroundContainer style={{flexDirection: "column", height: "fit-content", padding: "calc(10rem/16)"}}>
            <Text style={{width: "100%"}}>
                {label}
            </Text>
            <Text className="text" style={{width: "100%"}}>
                {text}
            </Text>
        </BackgroundContainer>
    )
}