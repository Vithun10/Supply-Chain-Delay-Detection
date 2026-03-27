import React, { useState, useEffect } from 'react';
import ShipmentTable from '../components/ShipmentTable';
import Filters from '../components/Filters';
import { fetchDelays, fetchHighRiskDelays } from '../services/api';
import { AlertTriangle, AlertOctagon } from 'lucide-react';

const DelaysPage = ({ type = 'all' }) => {
    const isHighRisk = type === 'high-risk';

    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const [page, setPage] = useState(1);
    const [filters, setFilters] = useState({});
    const [pagination, setPagination] = useState({ page: 1, limit: 50, totalPages: 1, totalRecords: 0 });

    const loadData = async (currentPage, currentFilters) => {
        try {
            setLoading(true);
            setError(null);

            const params = {
                page: currentPage,
                limit: 50,
                ...currentFilters
            };

            let response;
            if (isHighRisk) {
                // High risk endpoint currently only supports pagination filters in backend
                response = await fetchHighRiskDelays({ page: currentPage, limit: 50 });
            } else {
                response = await fetchDelays(params);
            }

            setData(response.data || []);
            setPagination({
                page: response.page,
                limit: response.limit,
                totalPages: response.total_pages,
                totalRecords: response.total_records
            });
        } catch (err) {
            console.error("Failed to fetch delays:", err);
            setError(`Failed to load ${isHighRisk ? 'high risk' : 'delayed'} shipments. Please check your connection.`);
        } finally {
            setLoading(false);
        }
    };

    // Reset page and filters when toggling between "Delayed" and "High Risk" via sidebar
    useEffect(() => {
        setPage(1);
        setFilters({});
    }, [type]);

    useEffect(() => {
        loadData(page, filters);
    }, [page, filters, type]);

    const handlePageChange = (newPage) => {
        setPage(newPage);
    };

    const handleFilterChange = (newFilters) => {
        setFilters(newFilters);
        setPage(1);
    };

    return (
        <div className="space-y-6">
            <div className="flex items-center space-x-3">
                {isHighRisk ? (
                    <div className="bg-red-100 p-2 rounded-lg text-red-600">
                        <AlertOctagon size={28} />
                    </div>
                ) : (
                    <div className="bg-amber-100 p-2 rounded-lg text-amber-600">
                        <AlertTriangle size={28} />
                    </div>
                )}
                <div>
                    <h2 className="text-2xl font-bold text-slate-800">
                        {isHighRisk ? 'High Risk Shipments' : 'Delayed Shipments'}
                    </h2>
                    <p className="text-slate-500">
                        {isHighRisk
                            ? 'Monitoring shipments with severe environmental or routing risks.'
                            : 'Review shipments that have missed their expected delivery windows.'}
                    </p>
                </div>
            </div>

            {!isHighRisk && (
                <Filters onFilterChange={handleFilterChange} showStatusFilter={false} />
            )}

            <ShipmentTable
                shipments={data}
                loading={loading}
                error={error}
                pagination={pagination}
                onPageChange={handlePageChange}
            />
        </div>
    );
};

export default DelaysPage;
