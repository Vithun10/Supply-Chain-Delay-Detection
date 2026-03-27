import React, { useState, useEffect } from 'react';
import AnalyticsCharts from '../components/AnalyticsCharts';
import { fetchTopDelayedRoutes, fetchCarrierPerformance, fetchAvgDeliveryTime } from '../services/api';
import { Loader2, AlertCircle } from 'lucide-react';

const AnalyticsPage = () => {
    const [chartsData, setChartsData] = useState({ routes: [], carriers: [], avgTimes: [] });
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const loadData = async () => {
            try {
                setLoading(true);
                setError(null);

                const [routesRef, carriersRef, avgTimesRef] = await Promise.all([
                    fetchTopDelayedRoutes(),
                    fetchCarrierPerformance(),
                    fetchAvgDeliveryTime()
                ]);

                setChartsData({
                    routes: routesRef.routes || [],
                    // Show top 10 for discrete full analytics page instead of top 5
                    carriers: carriersRef.carriers?.slice(0, 10) || [],
                    avgTimes: avgTimesRef.avg_delivery_times || []
                });

            } catch (err) {
                console.error("Analytics failed to load:", err);
                setError("Failed to load analytics data. Ensure the backend is running.");
            } finally {
                setLoading(false);
            }
        };

        loadData();
    }, []);

    return (
        <div className="space-y-6">
            <div>
                <h2 className="text-2xl font-bold text-slate-800">Deep Analytics</h2>
                <p className="text-slate-500">Comprehensive routing and carrier performance insights.</p>
            </div>

            {loading && (
                <div className="flex h-64 items-center justify-center bg-white rounded-xl border border-slate-200">
                    <Loader2 className="animate-spin text-primary-500 mb-4" size={48} />
                    <span className="ml-3 text-slate-600 font-medium">Processing Analytics...</span>
                </div>
            )}

            {error && !loading && (
                <div className="bg-red-50 p-6 rounded-lg border border-red-200 flex items-center shadow-sm">
                    <AlertCircle className="text-red-500 mr-4" size={28} />
                    <div>
                        <h3 className="text-red-800 font-bold text-lg">Error Loading Analytics</h3>
                        <p className="text-red-700 font-medium">{error}</p>
                    </div>
                </div>
            )}

            {!loading && !error && (
                <AnalyticsCharts data={chartsData} />
            )}
        </div>
    );
};

export default AnalyticsPage;
