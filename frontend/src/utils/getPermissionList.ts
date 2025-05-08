import {Permission} from "@/store/RolePageStore/store";

export default function getPermissionList() {
    const idPermission = [
        "canEditOrganization",
        "canDeleteOrganization",
        "canInviteEmployees",
        "canDeleteEmployees",
        "canEditProjects",
        "canEditTasks",
        "canEditSprints"
    ]
    const namePermission = [
        "Изменение данных организации",
        "Удаление организации",
        "Приглашение сотрудников",
        "Увольнение сотрудников",
        "Изменение проектов",
        "Изменение задач",
        "Изменение спринтов"
    ]
    const descriptionPermission = [
        "Добавляет возможность владельцу роли изменять данные организации",
        "Добавляет возможность владельцу роли удалять организацию",
        "Добавляет возможность владельцу роли приглашать сотрудников в организацию",
        "Добавляет возможность владельцу роли удалять пользователей из организации",
        "Добавляет возможность владельцу роли изменять данные проектов организации",
        "Добавляет возможность владельцу роли редактировать и создавать задачи",
        "Добавляет возможность владельцу роли редактировать спринт"
    ]
    var result:  Permission[] = []
    idPermission.map((role, index) => (
        result.push({id: idPermission[index], name: namePermission[index], description: descriptionPermission[index], state: false})
    ))
    return result;
}