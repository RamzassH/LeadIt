import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import {Paper, Typography} from "@mui/material";

interface BurndownChartProps {
    sprintDays: number;
    idealData: number[];
    actualData: number[];
}

const BurndownChart: React.FC<BurndownChartProps> = ({ sprintDays, idealData, actualData }) => {
    const data = Array.from({ length: sprintDays }, (_, i) => ({
        name: `Day ${i + 1}`,
        ideal: idealData[i],
        actual: actualData[i],
    }));

    return (
        <Paper sx={{ p: 2, mt: 3 }}>
            <Typography variant="h6" gutterBottom>
                Sprint Burndown
            </Typography>
            <ResponsiveContainer width="100%" height={300}>
                <LineChart data={data}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <Line type="monotone" dataKey="ideal" stroke="#ff0000" name="Ideal" />
                    <Line type="monotone" dataKey="actual" stroke="#00ff00" name="Actual" />
                </LineChart>
            </ResponsiveContainer>
        </Paper>
    );
};

export default BurndownChart;