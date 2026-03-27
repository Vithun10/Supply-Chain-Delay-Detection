import React from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';

const Pagination = ({ page, totalPages, totalRecords, limit, onPageChange }) => {
    const isFirst = page === 1;
    const isLast = page === totalPages || totalPages === 0;

    const startRecord = totalRecords === 0 ? 0 : (page - 1) * limit + 1;
    const endRecord = Math.min(page * limit, totalRecords);

    return (
        <div className="flex items-center justify-between px-4 py-3 bg-white border-t border-slate-200 sm:px-6">
            <div className="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
                <div>
                    <p className="text-sm text-slate-700">
                        Showing <span className="font-medium">{startRecord}</span> to <span className="font-medium">{endRecord}</span> of{' '}
                        <span className="font-medium">{totalRecords.toLocaleString()}</span> results
                    </p>
                </div>
                <div>
                    <nav className="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                        <button
                            onClick={() => onPageChange(page - 1)}
                            disabled={isFirst}
                            className={`relative inline-flex items-center px-2 py-2 rounded-l-md border border-slate-300 bg-white text-sm font-medium ${isFirst ? 'text-slate-300 cursor-not-allowed' : 'text-slate-500 hover:bg-slate-50'
                                }`}
                        >
                            <span className="sr-only">Previous</span>
                            <ChevronLeft size={20} />
                        </button>
                        <span className="relative inline-flex items-center px-4 py-2 border border-slate-300 bg-white text-sm font-medium text-slate-700">
                            Page {page} of {totalPages || 1}
                        </span>
                        <button
                            onClick={() => onPageChange(page + 1)}
                            disabled={isLast}
                            className={`relative inline-flex items-center px-2 py-2 rounded-r-md border border-slate-300 bg-white text-sm font-medium ${isLast ? 'text-slate-300 cursor-not-allowed' : 'text-slate-500 hover:bg-slate-50'
                                }`}
                        >
                            <span className="sr-only">Next</span>
                            <ChevronRight size={20} />
                        </button>
                    </nav>
                </div>
            </div>
        </div>
    );
};

export default Pagination;
