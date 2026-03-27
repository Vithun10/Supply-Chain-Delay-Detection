import React from 'react';
import { AlertCircle, FileX } from 'lucide-react';
import Pagination from './Pagination';

const ShipmentTable = ({ shipments, loading, error, pagination, onPageChange }) => {

    const formatDate = (dateString) => {
        if (!dateString) return 'N/A';
        return new Date(dateString).toLocaleDateString();
    };

    const getRiskColor = (level) => {
        switch (level) {
            case 'HIGH': return 'bg-red-100 text-red-800';
            case 'MEDIUM': return 'bg-amber-100 text-amber-800';
            case 'LOW': return 'bg-green-100 text-green-800';
            default: return 'bg-slate-100 text-slate-800';
        }
    };

    if (loading) {
        return (
            <div className="bg-white rounded-lg shadow-sm border border-slate-200 p-8 text-center">
                <div className="inline-block animate-spin rounded-full h-8 w-8 border-4 border-primary-500 border-t-transparent"></div>
                <p className="mt-2 text-slate-500 font-medium">Loading shipments...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div className="bg-red-50 p-6 rounded-lg border border-red-200 flex items-center shadow-sm">
                <AlertCircle className="text-red-500 mr-4" size={28} />
                <div>
                    <h3 className="text-red-800 font-bold text-lg">Error Loading Data</h3>
                    <p className="text-red-700 font-medium">{error}</p>
                </div>
            </div>
        );
    }

    if (!shipments || shipments.length === 0) {
        return (
            <div className="bg-white p-12 rounded-lg border border-slate-200 text-center shadow-sm">
                <FileX className="mx-auto text-slate-400 mb-4" size={48} />
                <h3 className="text-lg font-bold text-slate-800">No Shipments Found</h3>
                <p className="text-slate-500">Try adjusting your filters to see more results.</p>
            </div>
        );
    }

    return (
        <div className="bg-white shadow-sm rounded-lg border border-slate-200 overflow-hidden">
            <div className="overflow-x-auto">
                <table className="min-w-full divide-y divide-slate-200">
                    <thead className="bg-slate-50">
                        <tr>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">ID</th>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Route</th>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Carrier & Mode</th>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Exp. Delivery</th>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Delivered Date</th>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Delay</th>
                            <th scope="col" className="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Risk Level</th>
                        </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-slate-200">
                        {shipments.map((shipment) => (
                            <tr key={shipment.shipment_id} className="hover:bg-slate-50 transition-colors">
                                <td className="px-6 py-4 whitespace-nowrap text-sm font-bold text-slate-900">
                                    {shipment.shipment_id}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    <div className="text-sm text-slate-900 font-medium">{shipment.origin}</div>
                                    <div className="text-xs text-slate-500">&rarr; {shipment.destination}</div>
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    <div className="text-sm text-slate-900">{shipment.carrier}</div>
                                    <div className="text-xs text-slate-500">{shipment.mode}</div>
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
                                    {formatDate(shipment.expected_delivery_date)}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
                                    {formatDate(shipment.delivered_date)}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    {shipment.delay_detected ? (
                                        <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-bold bg-red-100 text-red-800">
                                            {Math.round(shipment.delay_days)} Days
                                        </span>
                                    ) : (
                                        <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-bold bg-green-100 text-green-800">
                                            On Time
                                        </span>
                                    )}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    <span className={`inline-flex items-center px-2.5 py-0.5 rounded font-bold text-xs ${getRiskColor(shipment.risk_level)}`}>
                                        {shipment.risk_level}
                                    </span>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>

            {pagination && (
                <Pagination
                    page={pagination.page}
                    limit={pagination.limit}
                    totalPages={pagination.totalPages}
                    totalRecords={pagination.totalRecords}
                    onPageChange={onPageChange}
                />
            )}
        </div>
    );
};

export default ShipmentTable;
