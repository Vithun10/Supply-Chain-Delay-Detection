import React, { useState } from 'react';
import { Search, MapPin, Calendar, AlertCircle } from 'lucide-react';
import { trackShipment } from '../services/api';

const CustomerTracking = () => {
    const [shipmentId, setShipmentId] = useState('');
    const [result, setResult] = useState(null);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);

    const handleSearch = async (e) => {
        e.preventDefault();
        if (!shipmentId.trim()) return;

        try {
            setLoading(true);
            setError(null);
            setResult(null);

            const data = await trackShipment(shipmentId.trim());
            setResult(data);
        } catch (err) {
            if (err.response && err.response.status === 404) {
                setError("Shipment not found. Please check your ID and try again.");
            } else {
                setError("Unable to retrieve tracking information at this time.");
            }
        } finally {
            setLoading(false);
        }
    };

    const formatDate = (dateString) => {
        if (!dateString) return 'N/A';
        return new Date(dateString).toLocaleDateString('en-US', {
            year: 'numeric', month: 'short', day: 'numeric'
        });
    };

    return (
        <div className="max-w-3xl mx-auto py-10">
            <div className="text-center mb-10">
                <h2 className="text-3xl font-bold text-slate-800">Track Your Shipment</h2>
                <p className="text-slate-500 mt-2">Enter your Shipment ID to get real-time delivery status.</p>
            </div>

            <div className="bg-white p-6 rounded-2xl shadow-sm border border-slate-200 mb-8">
                <form onSubmit={handleSearch} className="flex gap-4">
                    <div className="relative flex-1">
                        <Search className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400" size={20} />
                        <input
                            type="text"
                            value={shipmentId}
                            onChange={(e) => setShipmentId(e.target.value)}
                            placeholder="e.g. SHP100001"
                            className="w-full pl-12 pr-4 py-4 bg-slate-50 border border-slate-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-primary-500 outline-none transition-all font-medium text-slate-800"
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        disabled={loading}
                        className="bg-primary-600 hover:bg-primary-700 text-white px-8 py-4 rounded-xl font-bold transition-colors disabled:opacity-70 disabled:cursor-not-allowed"
                    >
                        {loading ? 'Searching...' : 'Track'}
                    </button>
                </form>
            </div>

            {error && (
                <div className="bg-red-50 border-l-4 border-red-500 p-4 rounded-r-lg">
                    <div className="flex items-center">
                        <AlertCircle className="text-red-500 mr-3" size={20} />
                        <p className="text-red-700 font-medium">{error}</p>
                    </div>
                </div>
            )}

            {result && (
                <div className="bg-white rounded-2xl shadow-sm border border-slate-200 overflow-hidden">
                    {/* Header */}
                    <div className={`p-6 text-white flex items-center justify-between ${result.delay_detected ? 'bg-amber-500' : 'bg-primary-600'
                        }`}>
                        <div>
                            <p className="text-white/80 text-sm font-semibold tracking-wider uppercase">Shipment Status</p>
                            <h3 className="text-2xl font-bold mt-1">
                                {result.delay_detected ? 'Delayed' : 'On Track'}
                            </h3>
                        </div>
                        <div className="text-right">
                            <p className="text-white/80 text-sm font-semibold tracking-wider uppercase">ID</p>
                            <p className="text-xl font-bold mt-1">{result.shipment_id}</p>
                        </div>
                    </div>

                    {/* Details */}
                    <div className="p-8">
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">

                            <div className="space-y-6">
                                <div className="flex items-start">
                                    <div className="bg-slate-100 p-3 rounded-full mr-4 text-slate-600">
                                        <MapPin size={24} />
                                    </div>
                                    <div>
                                        <p className="text-sm font-medium text-slate-500">Route</p>
                                        <p className="text-lg font-bold text-slate-800">{result.origin} &rarr; {result.destination}</p>
                                        <p className="text-sm text-slate-500 capitalize mt-1">{result.mode} transit via {result.carrier}</p>
                                    </div>
                                </div>

                                <div className="flex items-start">
                                    <div className="bg-slate-100 p-3 rounded-full mr-4 text-slate-600">
                                        <Calendar size={24} />
                                    </div>
                                    <div>
                                        <p className="text-sm font-medium text-slate-500">Expected Delivery</p>
                                        <p className="text-lg font-bold text-slate-800">{formatDate(result.expected_delivery_date)}</p>
                                        {result.delay_detected && (
                                            <p className="text-sm text-red-600 font-medium mt-1">
                                                Running {Math.round(result.delay_days)} days behind schedule
                                            </p>
                                        )}
                                    </div>
                                </div>
                            </div>

                            <div className="bg-slate-50 p-6 rounded-xl border border-slate-200">
                                <h4 className="font-bold text-slate-800 mb-4 border-b border-slate-200 pb-2">Analysis Snapshot</h4>
                                <div className="space-y-3">
                                    <div className="flex justify-between">
                                        <span className="text-slate-500 font-medium">Risk Level</span>
                                        <span className={`font-bold px-2 py-0.5 rounded text-xs leading-tight ${result.risk_level === 'HIGH' ? 'bg-red-100 text-red-700' :
                                                result.risk_level === 'MEDIUM' ? 'bg-amber-100 text-amber-700' :
                                                    'bg-green-100 text-green-700'
                                            }`}>
                                            {result.risk_level}
                                        </span>
                                    </div>
                                    <div className="flex justify-between">
                                        <span className="text-slate-500 font-medium">Delay Detected</span>
                                        <span className="text-slate-800 font-semibold">{result.delay_detected ? 'Yes' : 'No'}</span>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default CustomerTracking;
