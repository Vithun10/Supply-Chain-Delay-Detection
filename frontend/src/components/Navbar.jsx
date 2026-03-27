import { Boxes } from 'lucide-react';

const Navbar = ({ role, setRole }) => {
    return (
        <nav className="bg-white border-b border-slate-200 px-6 py-3 flex items-center justify-between shadow-sm z-10 relative">
            <div className="flex items-center space-x-3">
                <div className="bg-primary-500 p-2 rounded-lg text-white">
                    <Boxes size={24} />
                </div>
                <div>
                    <h1 className="text-xl font-bold text-slate-800 leading-tight">SupplyChain<span className="text-primary-600">Sync</span></h1>
                    <p className="text-xs text-slate-500 font-medium tracking-wide">Delay Detection & Monitoring</p>
                </div>
            </div>

            <div className="flex items-center space-x-4">
                <label htmlFor="role-select" className="text-sm font-medium text-slate-600">
                    Viewing as:
                </label>
                <select
                    id="role-select"
                    value={role}
                    onChange={(e) => setRole(e.target.value)}
                    className="bg-slate-50 border border-slate-300 text-slate-800 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block p-2 font-medium"
                >
                    <option value="Owner">Owner Role</option>
                    <option value="Admin">Admin Role</option>
                    <option value="Customer">Customer Role</option>
                </select>
            </div>
        </nav>
    );
};

export default Navbar;
