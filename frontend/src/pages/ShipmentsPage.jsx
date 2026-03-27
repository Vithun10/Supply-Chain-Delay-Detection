import React, { useState, useEffect } from 'react';
import ShipmentTable from '../components/ShipmentTable';
import Filters from '../components/Filters';
import { fetchShipments } from '../services/api';

const ShipmentsPage = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const [page, setPage] = useState(1);
    const [filters, setFilters] = useState({});
    const [pagination, setPagination] = useState({ page: 1, limit: 50, totalPages: 1, totalRecords: 0 });

    const loadShipments = async (currentPage, currentFilters) => {
        try {
            setLoading(true);
            setError(null);

            const params = {
                page: currentPage,
                limit: 50,
                ...currentFilters
            };

            const response = await fetchShipments(params);

            setData(response.data || []);
            setPagination({
                page: response.page,
                limit: response.limit,
                totalPages: response.total_pages,
                totalRecords: response.total_records
            });
        } catch (err) {
            console.error("Failed to fetch shipments:", err);
            setError("Failed to load shipments. Please check your connection to the server.");
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadShipments(page, filters);
    }, [page, filters]);

    const handlePageChange = (newPage) => {
        setPage(newPage);
    };

    const handleFilterChange = (newFilters) => {
        setFilters(newFilters);
        setPage(1); // Reset to first page on new filter
    };

    return (
        <div className="space-y-6">
            <div>
                <h2 className="text-2xl font-bold text-slate-800">All Shipments</h2>
                <p className="text-slate-500">View and filter the complete logistics database.</p>
            </div>

            <Filters onFilterChange={handleFilterChange} showStatusFilter={true} />

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

export default ShipmentsPage;
