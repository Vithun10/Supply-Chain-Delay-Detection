import React, { useState } from 'react';
import { Filter, X } from 'lucide-react';

const Filters = ({ onFilterChange, showStatusFilter = true }) => {
    const [filters, setFilters] = useState({
        origin: '',
        destination: '',
        carrier: '',
        mode: '',
        status: ''
    });

    const [isOpen, setIsOpen] = useState(false);

    const handleChange = (e) => {
        const { name, value } = e.target;
        const newFilters = { ...filters, [name]: value };
        setFilters(newFilters);
    };

    const applyFilters = (e) => {
        e.preventDefault();
        onFilterChange(filters);
    };

    const clearFilters = () => {
        const empty = { origin: '', destination: '', carrier: '', mode: '', status: '' };
        setFilters(empty);
        onFilterChange(empty);
    };

    return (
        <div className="bg-white rounded-lg border border-slate-200 shadow-sm mb-6">
            <button
                onClick={() => setIsOpen(!isOpen)}
                className="w-full flex items-center justify-between p-4 focus:outline-none hover:bg-slate-50 rounded-lg transition-colors"
            >
                <div className="flex items-center space-x-2 text-slate-700 font-medium">
                    <Filter size={20} />
                    <span>Filter Shipments</span>
                </div>
                <span className="text-primary-600 text-sm font-medium">
                    {isOpen ? 'Close' : 'Expand'}
                </span>
            </button>

            {isOpen && (
                <form onSubmit={applyFilters} className="p-4 border-t border-slate-200 bg-slate-50 rounded-b-lg">
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4">

                        <div>
                            <label className="block text-xs font-semibold text-slate-600 uppercase mb-1">Origin</label>
                            <input
                                type="text"
                                name="origin"
                                value={filters.origin}
                                onChange={handleChange}
                                placeholder="e.g. Mumbai"
                                className="w-full border border-slate-300 rounded p-2 text-sm focus:ring-primary-500 focus:border-primary-500"
                            />
                        </div>

                        <div>
                            <label className="block text-xs font-semibold text-slate-600 uppercase mb-1">Destination</label>
                            <input
                                type="text"
                                name="destination"
                                value={filters.destination}
                                onChange={handleChange}
                                placeholder="e.g. Delhi"
                                className="w-full border border-slate-300 rounded p-2 text-sm focus:ring-primary-500 focus:border-primary-500"
                            />
                        </div>

                        <div>
                            <label className="block text-xs font-semibold text-slate-600 uppercase mb-1">Carrier</label>
                            <input
                                type="text"
                                name="carrier"
                                value={filters.carrier}
                                onChange={handleChange}
                                placeholder="e.g. BlueDart"
                                className="w-full border border-slate-300 rounded p-2 text-sm focus:ring-primary-500 focus:border-primary-500"
                            />
                        </div>

                        <div>
                            <label className="block text-xs font-semibold text-slate-600 uppercase mb-1">Mode</label>
                            <select
                                name="mode"
                                value={filters.mode}
                                onChange={handleChange}
                                className="w-full border border-slate-300 rounded p-2 text-sm focus:ring-primary-500 focus:border-primary-500 bg-white"
                            >
                                <option value="">All Modes</option>
                                <option value="Road">Road</option>
                                <option value="Air">Air</option>
                                <option value="Ocean">Ocean</option>
                                <option value="Rail">Rail</option>
                            </select>
                        </div>

                        {showStatusFilter && (
                            <div>
                                <label className="block text-xs font-semibold text-slate-600 uppercase mb-1">Delay Status</label>
                                <select
                                    name="status"
                                    value={filters.status}
                                    onChange={handleChange}
                                    className="w-full border border-slate-300 rounded p-2 text-sm focus:ring-primary-500 focus:border-primary-500 bg-white"
                                >
                                    <option value="">All Statuses</option>
                                    <option value="delayed">Delayed</option>
                                    <option value="ontime">On Time</option>
                                </select>
                            </div>
                        )}
                    </div>

                    <div className="mt-4 flex justify-end space-x-3">
                        <button
                            type="button"
                            onClick={clearFilters}
                            className="flex items-center px-4 py-2 border border-slate-300 shadow-sm text-sm font-medium rounded-md text-slate-700 bg-white hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
                        >
                            <X size={16} className="mr-2" />
                            Clear
                        </button>
                        <button
                            type="submit"
                            className="flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
                        >
                            <Filter size={16} className="mr-2" />
                            Apply Filters
                        </button>
                    </div>
                </form>
            )}
        </div>
    );
};

export default Filters;
