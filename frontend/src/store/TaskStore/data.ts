import {ProgressBar, Status, Task} from "@/store/TaskStore/store";

export const progressBar: ProgressBar = {
    id: 1,
    project_id: 1,
    name: "Гойда"
}

export const status: Status[] = [
    {
        id: 1,
        name: "To do",
        order: 1,
        notify_roles_ids: [],
        progress_bar_id: 1
    },
    {
        id: 2,
        name: "In Progress",
        order: 2,
        notify_roles_ids: [],
        progress_bar_id: 1
    },
    {
        id: 3,
        name: "Done",
        order: 3,
        notify_roles_ids: [],
        progress_bar_id: 1
    },
]

export const tasks: Task[] = [
    {
        id: 1,
        name: "Написать API",
        description: "Написать интерфейс, реализацию для запросов через API Gateway",
        project_id: 1,
        setter_id:  1,
        solver_id: 1,
        progress_bar_id: 1,
        current_status_id: 1,
        allocated_time: 5,
        wasted_time: 0,
        active: false,
        last_start_time: 0,
    },
    {
        id: 2,
        name: "Рефакторинг кода",
        description: "Поправить код страницы организаций",
        project_id: 1,
        setter_id:  1,
        solver_id: 1,
        progress_bar_id: 1,
        current_status_id: 2,
        allocated_time: 5,
        wasted_time: 1,
        active: true,
        last_start_time: 3,
    },
    {
        id: 3,
        name: "Исправить разметку профиля пользователя",
        description: "Убрать неиспользуемые поля",
        project_id: 1,
        setter_id:  1,
        solver_id: 1,
        progress_bar_id: 1,
        current_status_id: 3,
        allocated_time: 5,
        wasted_time: 0,
        active: false,
        last_start_time: 0,
    },
    {
        id: 4,
        name: "Изменить дизайн главной страницы",
        description: "Добавить темную тему",
        project_id: 1,
        setter_id:  1,
        solver_id: 1,
        progress_bar_id: 1,
        current_status_id: 3,
        allocated_time: 5,
        wasted_time: 0,
        active: false,
        last_start_time: 0,
    },
    {
        id: 5,
        name: "Добавить страницу проекта",
        description: "Написать разметку страницы проекта по спецификации",
        project_id: 1,
        setter_id:  1,
        solver_id: 1,
        progress_bar_id: 1,
        current_status_id: 1,
        allocated_time: 5,
        wasted_time: 0,
        active: false,
        last_start_time: 0,
    },
    {
        id: 6,
        name: "Написать нормальный текст отчета для диплома",
        description: "Мне хочется плакать...",
        project_id: 1,
        setter_id:  1,
        solver_id: 1,
        progress_bar_id: 1,
        current_status_id: 5,
        allocated_time: 5,
        wasted_time: 0,
        active: false,
        last_start_time: 0,
    }
]