import React from 'react';
import {
    BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip as RechartsTooltip, ResponsiveContainer,
    PieChart, Pie, Cell, Legend
} from 'recharts';

const COLORS = ['#14b8a6', '#f59e0b', '#ef4444', '#8b5cf6', '#3b82f6'];

const AnalyticsCharts = ({ data }) => {
    const { routes, carriers, avgTimes } = data;

    // Format data for Recharts
    const routeData = routes?.map(r => ({
        name: `${r.origin} \u2192 ${r.destination}`,
        delays: r.delay_count
    })) || [];

    const carrierData = carriers?.slice(0, 5).map(c => ({
        name: c.carrier,
        performance: (100 - c.delay_rate_percent).toFixed(1)
    })) || [];

    const delayRateData = carriers?.slice(0, 5).map(c => ({
        name: c.carrier,
        rate: c.delay_rate_percent.toFixed(1)
    })) || [];

    return (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">

            {/* Top Delayed Routes */}
            <div className="bg-white p-5 rounded-xl border border-slate-200 shadow-sm">
                <h3 className="text-lg font-bold text-slate-800 mb-4">Top Delayed Routes</h3>
                <div className="h-72">
                    <ResponsiveContainer width="100%" height="100%">
                        <BarChart data={routeData} layout="vertical" margin={{ top: 5, right: 30, left: 60, bottom: 5 }}>
                            <CartesianGrid strokeDasharray="3 3" horizontal={false} stroke="#e2e8f0" />
                            <XAxis type="number" textAnchor="end" />
                            <YAxis dataKey="name" type="category" width={120} tick={{ fontSize: 12 }} />
                            <RechartsTooltip cursor={{ fill: '#f8fafc' }} />
                            <Bar dataKey="delays" fill="#f59e0b" radius={[0, 4, 4, 0]} name="Delay Count" />
                        </BarChart>
                    </ResponsiveContainer>
                </div>
            </div>

            {/* Carrier On-Time Performance */}
            <div className="bg-white p-5 rounded-xl border border-slate-200 shadow-sm">
                <h3 className="text-lg font-bold text-slate-800 mb-4">Top Carriers (On-Time %)</h3>
                <div className="h-72">
                    <ResponsiveContainer width="100%" height="100%">
                        <BarChart data={carrierData} margin={{ top: 5, right: 30, left: 0, bottom: 5 }}>
                            <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#e2e8f0" />
                            <XAxis dataKey="name" tick={{ fontSize: 12 }} />
                            <YAxis domain={[0, 100]} />
                            <RechartsTooltip cursor={{ fill: '#f8fafc' }} />
                            <Bar dataKey="performance" fill="#14b8a6" radius={[4, 4, 0, 0]} name="On-Time %" />
                        </BarChart>
                    </ResponsiveContainer>
                </div>
            </div>

            {/* Carrier Delay Rate Comparison */}
            <div className="bg-white p-5 rounded-xl border border-slate-200 shadow-sm lg:col-span-2">
                <h3 className="text-lg font-bold text-slate-800 mb-4">Carrier Delay Rate Comparison</h3>
                <div className="h-80 flex items-center justify-center">
                    <ResponsiveContainer width="100%" height="100%">
                        <PieChart>
                            <Pie
                                data={delayRateData}
                                cx="50%"
                                cy="50%"
                                innerRadius={80}
                                outerRadius={120}
                                paddingAngle={5}
                                dataKey="rate"
                                nameKey="name"
                                label={({ name, percent }) => `${name} ${(percent * 100).toFixed(0)}%`}
                            >
                                {delayRateData.map((entry, index) => (
                                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                                ))}
                            </Pie>
                            <RechartsTooltip formatter={(value) => `${value}%`} />
                            <Legend verticalAlign="bottom" height={36} />
                        </PieChart>
                    </ResponsiveContainer>
                </div>
            </div>

        </div>
    );
};

export default AnalyticsCharts;
