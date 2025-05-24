"use client";
import {
    DrawerContainer,
    DrawerBackground,
    DrawerLogo
} from "@/components/UI/DrawerSideMenu/styled/DrawerContainer";
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import Icon from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/StyleIcon";
import VIPIcon from "@/images/icons-svg/VIPIcon";
import Text from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/Text";
import MenuDropList from "@/components/UI/ProfilePage/SideMenu/MenuDropList/MenuDropList";
import React, { forwardRef, useImperativeHandle, useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import CreateOrganizationModalWindow
    , {
    CreateOrganizationForm
} from "@/components/UI/ModalWindowTemplate/CreateOrganizationModalWindow/CreateOrganizationModalWindow";
import {useFetching} from "@/hooks/useFetching";
import {getOrganizationsAPI} from "@/api/organization/get";
import {createOrganizationAPI} from "@/api/organization/create";
import useOrganizationStore, {Organization} from "@/store/OrganizationStore/store";
import useGlobalStore from "@/store/GlobalStore/store";
import Image from "next/image";
import useProjectStore from "@/store/ProjectsPageStore/store";

interface ButtonContent {
    icon: React.ReactNode;
    text: string;
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

interface DrawerSideMenuProps {
    // В родительском компоненте будет вызываться handleClick через ref
}

interface DrawerSideMenuRef {
    triggerHandleClick: () => void;
}

interface OrganizationResponseForm {
    id: number,
    name: string,
    organizer_id: number,
    description: string,
    image: string,
}

const DrawerSideMenu = forwardRef<DrawerSideMenuRef, DrawerSideMenuProps>((props, ref) => {
    const {organizations, setOrganizations} = useOrganizationStore();
    const currentOrganization = useGlobalStore(state => state.currentOrganization);
    const setCurrentOrganization = useGlobalStore(state => state.setCurrentOrganization);
    const router = useRouter();
    const [open, setOpen] = useState(false);
    const [openCreateWindow, setOpenCreateWindow] = useState(false);
    const [createOrganizationRequest, isLoadingCreateRequest, errorCreateRequest] = useFetching(async (data: CreateOrganizationForm) => {
        const response = await createOrganizationAPI(data, "")
        const organization = {...data, organizer_id: 1, id: response.data};
        setCurrentOrganization(organization);
        setOrganizations(
            [organization, ...organizations]
        )
    })
    const [getOrganizationsList, isLoadingOrganizationsList, errorOrganizationsList] = useFetching(async () => {
        const response = await getOrganizationsAPI(1, "");
        setOrganizations(response.data);
    })
    useEffect(() => {
        getOrganizationsList();
    }, []);

    const create = async (data: CreateOrganizationForm) => {
        await createOrganizationRequest(data)
        setOpen(false);
    }

    const openModalWindowCreateOrganization = (event: React.MouseEvent<HTMLButtonElement>) => {
        setOpenCreateWindow(true);
    }
    const toggleDrawer = (newOpen: boolean) => {
        setOpen(newOpen);
    };
    useImperativeHandle(ref, () => ({
        triggerHandleClick: () => { toggleDrawer(true); }
    }));

    const list1: ButtonContent[] = [
        { icon: <VIPIcon />, text: "Главная", callback: (event) => {router?.push("/organization")} },
        { icon: <VIPIcon />, text: "Структура", callback: (event) => {router?.push("/organization/structure")} },
        { icon: <VIPIcon />, text: "Роли", callback: (event) => {router?.push("/organization/roles")} },
        { icon: <VIPIcon />, text: "Проекты", callback: (event) => {router?.push("/organization/projects")} },
    ];

    return (
        <DrawerContainer open={open} onClose={() => { toggleDrawer(false); }}>
            <DrawerBackground>
                <DrawerLogo>
                    LeadIt
                </DrawerLogo>
                <MenuButton callback={(event) => { router?.push("/"); }}>
                    <Icon className="first-icon">
                        <VIPIcon />
                    </Icon>
                    <Text>Главная</Text>
                </MenuButton>
                <MenuButton callback={(event) => { router?.push("/profile"); }}>
                    <Icon className="first-icon">
                        <VIPIcon/>
                    </Icon>
                    <Text>Профиль</Text>
                </MenuButton>
                {
                    organizations.length < 1?
                        null:
                        <MenuDropList firstIcon={<VIPIcon/>} text={currentOrganization? currentOrganization.name: "Организация"}>
                            {list1.map((item, index) => (
                                <MenuButton className="list-item" callback={item.callback} key={index}>
                                    <Icon className="first-icon">{item.icon}</Icon>
                                    <Text>{item.text}</Text>
                                </MenuButton>
                            ))}
                        </MenuDropList>
                }
                <MenuDropList firstIcon={<VIPIcon/>} text="Список организаций">
                    {organizations.map((item, index) => (
                        <MenuButton className="list-item" callback={() => {setCurrentOrganization(item)}} key={index}>
                            <Icon className="first-icon"><Image src={item.image} alt="Ошибочка вышла"/></Icon>
                            <Text>{item.name}</Text>
                        </MenuButton>
                    ))}
                    <MenuButton className="list-item" callback={openModalWindowCreateOrganization} key={organizations.length}>
                        <Text>Создать организацию</Text>
                    </MenuButton>
                </MenuDropList>
            </DrawerBackground>
            <CreateOrganizationModalWindow callback={create} open={openCreateWindow} handleClose={() => {setOpenCreateWindow(false);}}/>
        </DrawerContainer>
    );
});

export default DrawerSideMenu;