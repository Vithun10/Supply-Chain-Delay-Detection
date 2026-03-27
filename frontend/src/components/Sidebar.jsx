import { Link, useLocation } from 'react-router-dom';
import { LayoutDashboard, Truck, AlertTriangle, Activity, BarChart3, Upload } from 'lucide-react';

const Sidebar = ({ role }) => {
    const location = useLocation();

    const menuItems = [
        { name: 'Dashboard', path: '/', icon: LayoutDashboard, roles: ['Owner', 'Admin'] },
        { name: 'Shipments', path: '/shipments', icon: Truck, roles: ['Owner', 'Admin'] },
        { name: 'Delayed Shipments', path: '/delays', icon: AlertTriangle, roles: ['Owner', 'Admin'] },
        { name: 'High Risk', path: '/delays/high-risk', icon: Activity, roles: ['Owner', 'Admin'] },
        { name: 'Analytics', path: '/analytics', icon: BarChart3, roles: ['Owner'] },

        // NEW ADMIN FEATURE
        { name: 'Upload Dataset', path: '/upload-dataset', icon: Upload, roles: ['Admin'] },
    ];

    const visibleItems = menuItems.filter(item => item.roles.includes(role));

    return (
        <div className="w-64 bg-white border-r border-slate-200 shadow-sm flex flex-col h-full z-0">
            <div className="flex-1 py-6 px-4 space-y-2">
                {visibleItems.map((item) => {
                    const Icon = item.icon;
                    const isActive =
                        location.pathname === item.path ||
                        (location.pathname.startsWith(item.path) && item.path !== '/');

                    return (
                        <Link
                            key={item.name}
                            to={item.path}
                            className={`flex items-center space-x-3 px-4 py-3 rounded-lg font-medium transition-colors ${isActive
                                    ? 'bg-primary-50 text-primary-700'
                                    : 'text-slate-600 hover:bg-slate-50 hover:text-slate-900'
                                }`}
                        >
                            <Icon size={20} className={isActive ? 'text-primary-600' : 'text-slate-400'} />
                            <span>{item.name}</span>
                        </Link>
                    );
                })}
            </div>
        </div>
    );
};

export default Sidebar;