import React, { useState, useEffect } from 'react';
import DashboardCards from '../components/DashboardCards';
import AnalyticsCharts from '../components/AnalyticsCharts';
import { fetchDelayRate, fetchTopDelayedRoutes, fetchCarrierPerformance, fetchAvgDeliveryTime, fetchHighRiskDelays, fetchShipments } from '../services/api';
import { Loader2 } from 'lucide-react';

const OwnerDashboard = () => {
    const [metrics, setMetrics] = useState({ total: 0, delayed: 0, highRisk: 0, delayRate: 0 });
    const [chartsData, setChartsData] = useState({ routes: [], carriers: [], avgTimes: [] });
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const loadData = async () => {
            try {
                setLoading(true);
                // We get total and delay rate from delay-rate endpoint
                const rateData = await fetchDelayRate();

                // We get high risk count by fetching the first page of high-risk endpoint
                const hrData = await fetchHighRiskDelays({ page: 1, limit: 1 });

                setMetrics({
                    total: rateData.total_shipments || 0,
                    delayed: rateData.delayed_shipments || 0,
                    highRisk: hrData.total_records || 0,
                    delayRate: rateData.delay_rate_percent || 0
                });

                // Fetch charts data
                const routesRef = await fetchTopDelayedRoutes();
                const carriersRef = await fetchCarrierPerformance();
                const avgTimesRef = await fetchAvgDeliveryTime();

                setChartsData({
                    routes: routesRef.routes.slice(0, 5), // top 5 routes for standard dashboard view
                    carriers: carriersRef.carriers,
                    avgTimes: avgTimesRef.avg_delivery_times
                });

            } catch (err) {
                console.error("Dashboard failed to load:", err);
                setError("Failed to load dashboard data. Ensure the backend is running.");
            } finally {
                setLoading(false);
            }
        };

        loadData();
    }, []);

    if (loading) {
        return (
            <div className="flex h-full items-center justify-center">
                <Loader2 className="animate-spin text-primary-500 mb-4" size={48} />
                <span className="ml-3 text-slate-600 font-medium">Loading Dashboard...</span>
            </div>
        );
    }

    if (error) {
        return (
            <div className="bg-red-50 text-red-600 p-4 rounded-lg border border-red-200">
                <p className="font-medium">{error}</p>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            <div>
                <h2 className="text-2xl font-bold text-slate-800">Owner Dashboard</h2>
                <p className="text-slate-500">System overview and analytics at a glance.</p>
            </div>

            <DashboardCards metrics={metrics} />

            <div className="mt-8">
                <AnalyticsCharts data={chartsData} />
            </div>
        </div>
    );
};

export default OwnerDashboard;
