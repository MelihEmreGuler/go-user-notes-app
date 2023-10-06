import Layout from "../layouts/Layout.jsx";
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';


export default function Test(){

    const data = [
        {
            name: 'Mon',
            count: 3,
        },
        {
            name: 'Tue',
            count: 10,
        },
        {
            name: 'Wed',
            count: 15,
        },
        {
            name: 'Thu',
            count: 30,
        },
        {
            name: 'Fri',
            count: 5,
        },
        {
            name: 'Sat',
            count: 50,
        },
        {
            name: 'Sun',
            count: 12,
        },
    ];

    return(
        <Layout>
            <ResponsiveContainer width="50%" height="50%">
                <BarChart
                    width={500}
                    height={300}
                    data={data}
                    margin={{
                        top: 5,
                        right: 30,
                        left: 20,
                        bottom: 5,
                    }}
                    barSize={20}
                >
                    <XAxis dataKey="name" scale="point" padding={{ left: 5, right: 5 }} />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <CartesianGrid strokeDasharray="3 3" />
                    <Bar dataKey="count" fill="#8884d8" background={{ fill: '#eee' }} />
                </BarChart>
            </ResponsiveContainer>
        </Layout>
    )
}