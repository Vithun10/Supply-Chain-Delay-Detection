import { Users, Truck, AlertTriangle, AlertOctagon } from 'lucide-react';

const DashboardCards = ({ metrics }) => {
    const { total, delayed, highRisk, delayRate } = metrics;

    const cards = [
        {
            title: 'Total Shipments',
            value: total.toLocaleString(),
            icon: Truck,
            color: 'bg-blue-100 text-blue-600',
        },
        {
            title: 'Delayed Shipments',
            value: delayed.toLocaleString(),
            icon: AlertTriangle,
            color: 'bg-amber-100 text-amber-600',
        },
        {
            title: 'High Risk Shipments',
            value: highRisk.toLocaleString(),
            icon: AlertOctagon,
            color: 'bg-red-100 text-red-600',
        },
        {
            title: 'Delay Rate',
            value: `${delayRate.toFixed(1)}%`,
            icon: Users, // Using a generic placeholder icon structure
            color: 'bg-purple-100 text-purple-600',
        }
    ];

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
            {cards.map((card, idx) => {
                const Icon = card.icon;
                return (
                    <div key={idx} className="bg-white rounded-xl border border-slate-200 p-6 flex items-center shadow-sm">
                        <div className={`p-4 rounded-full ${card.color} mr-4`}>
                            <Icon size={24} />
                        </div>
                        <div>
                            <p className="text-sm font-medium text-slate-500 mb-1">{card.title}</p>
                            <h3 className="text-2xl font-bold text-slate-800">{card.value}</h3>
                        </div>
                    </div>
                );
            })}
        </div>
    );
};

export default DashboardCards;
