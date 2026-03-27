import React from 'react';
import { Truck, ShieldCheck, Activity } from 'lucide-react';
import { Link } from 'react-router-dom';

const AdminDashboard = () => {
    return (
        <div className="space-y-6">
            <div>
                <h2 className="text-2xl font-bold text-slate-800">Admin Workspace</h2>
                <p className="text-slate-500">Manage and oversee shipment records.</p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <Link to="/shipments" className="bg-white p-6 rounded-xl border border-slate-200 shadow-sm hover:shadow-md transition-shadow group">
                    <div className="bg-blue-100 w-12 h-12 rounded-lg flex items-center justify-center text-blue-600 mb-4 group-hover:bg-blue-600 group-hover:text-white transition-colors">
                        <Truck size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-slate-800">All Shipments</h3>
                    <p className="text-sm text-slate-500 mt-2">View and filter through the complete database of logistics records.</p>
                </Link>

                <Link to="/delays" className="bg-white p-6 rounded-xl border border-slate-200 shadow-sm hover:shadow-md transition-shadow group">
                    <div className="bg-amber-100 w-12 h-12 rounded-lg flex items-center justify-center text-amber-600 mb-4 group-hover:bg-amber-500 group-hover:text-white transition-colors">
                        <ShieldCheck size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-slate-800">Delayed Shipments</h3>
                    <p className="text-sm text-slate-500 mt-2">Investigate shipments that missed their expected delivery dates.</p>
                </Link>

                <Link to="/delays/high-risk" className="bg-white p-6 rounded-xl border border-slate-200 shadow-sm hover:shadow-md transition-shadow group">
                    <div className="bg-red-100 w-12 h-12 rounded-lg flex items-center justify-center text-red-600 mb-4 group-hover:bg-red-600 group-hover:text-white transition-colors">
                        <Activity size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-slate-800">High Risk</h3>
                    <p className="text-sm text-slate-500 mt-2">Monitor shipments flagged with severe environmental or operational risks.</p>
                </Link>
            </div>
        </div>
    );
};

export default AdminDashboard;
