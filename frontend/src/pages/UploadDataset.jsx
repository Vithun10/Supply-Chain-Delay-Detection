import { useState } from "react";
import axios from "axios";
import { UploadCloud } from "lucide-react";
import { useNavigate } from "react-router-dom";

const UploadDataset = () => {
    const [file, setFile] = useState(null);
    const [message, setMessage] = useState("");
    const [uploading, setUploading] = useState(false);

    const navigate = useNavigate();

    const handleFileChange = (e) => {
        setFile(e.target.files[0]);
        setMessage("");
    };

    const handleUpload = async () => {
        if (!file) {
            setMessage("Please select a CSV file.");
            return;
        }

        const formData = new FormData();
        formData.append("file", file);

        try {
            setUploading(true);

            const response = await axios.post(
                "http://localhost:8080/upload-dataset",
                formData,
                {
                    headers: {
                        "Content-Type": "multipart/form-data",
                    },
                }
            );

            setMessage(`✅ ${response.data.message}`);

            setTimeout(() => {
                navigate("/");
                window.location.reload();
            }, 2000);

        } catch (error) {
            setMessage("❌ Upload failed. Please check dataset format.");
            setUploading(false);
        }
    };

    return (
        <div className="p-8 max-w-xl mx-auto">
            <div className="bg-white shadow-lg rounded-xl p-6 border border-slate-200">

                <h1 className="text-2xl font-bold mb-4 flex items-center gap-2">
                    <UploadCloud size={24} />
                    Upload New Dataset
                </h1>

                <p className="text-slate-500 mb-6">
                    Upload a CSV dataset to reprocess shipments and update analytics.
                </p>

                <input
                    type="file"
                    accept=".csv"
                    onChange={handleFileChange}
                    className="mb-4 block w-full text-sm"
                />

                <button
                    onClick={handleUpload}
                    disabled={uploading}
                    className="bg-primary-600 text-white px-6 py-2 rounded-lg hover:bg-primary-700 transition"
                >
                    {uploading ? "Uploading..." : "Upload Dataset"}
                </button>

                {message && (
                    <p className="mt-4 text-sm font-medium text-slate-700">
                        {message}
                    </p>
                )}
            </div>
        </div>
    );
};

export default UploadDataset;