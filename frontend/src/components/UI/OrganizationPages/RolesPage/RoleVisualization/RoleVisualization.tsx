import React from "react";
import {
    Header, RoleNameInput,
    VisualizationCard, VisualizationDescription,
    VisualizationItem
} from "@/components/UI/OrganizationPages/RolesPage/RoleVisualization/styled/styled";
import {Box, Button, Divider, Switch, Typography} from "@mui/material";

interface RoleVisualizationProps {

}

const RoleVisualization: React.FC<RoleVisualizationProps> = ({
                                               }) => {
    const [roleName, setRoleName] = React.useState<string>();
    const callback = () => {

    }

    return (
        <VisualizationCard>
            <Header>
                <Typography variant="h6">Настройки роли</Typography>
                <Button
                    variant="contained"
                    color="secondary"
                    onClick={callback}
                >
                    Сохранить изменения
                </Button>
            </Header>

            <Divider />


            <React.Fragment key="Наименование роли">
                <VisualizationItem>
                    <Box>
                        <RoleNameInput
                            label="Название роли"
                            value={roleName}
                            onChange={(e) => setRoleName(e.target.value)}
                            required
                            fullWidth
                            variant="outlined"
                        />
                    </Box>
                </VisualizationItem>
            </React.Fragment>
            <Divider/>
        </VisualizationCard>
    );
};

export default RoleVisualization;