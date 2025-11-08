import { useEffect, useState } from "react";
import axios from "axios";
import { useGlobal } from "./GlobalState";

interface UrlItem {
  ID: number;
  Url: string;
  ShortUrl: string;
  UserID: string;
}

export default function UrlList() {
  const [urls, setUrls] = useState<UrlItem[]>([]);
  const { toggle, setToggle } = useGlobal();

  useEffect(() => {
    const fetchUrls = async () => {
      try {
        const res = await axios.get<UrlItem[]>("http://13.213.2.17:8080/urls");
        console.log(res.data);
        setUrls(res.data);
      } catch (err) {
        alert("Failed to fetch URLs " + err);
      }
    };
    fetchUrls();
  }, [toggle]);

  const handleDelete = async (name: string) => {
    if (!confirm(`Delete ${name}?`)) return;
    try {
      await axios.delete(`http://13.213.2.17:8080/${name}`);
      setToggle(!toggle);
    } catch (err) {
      alert("Failed to delete: " + err);
    }
  };

  return (
    <div className="mt-8 bg-white p-6 rounded-2xl shadow-lg w-full max-w-3xl mx-auto">
      <h2 className="text-xl font-semibold text-gray-700 mb-4">URL List</h2>

      <table className="w-full border-collapse border border-gray-300 text-sm">
        <thead className="bg-gray-100">
          <tr>
            <th className="border px-4 py-2 text-left">#</th>
            <th className="border px-4 py-2 text-left">Original URL</th>
            <th className="border px-4 py-2 text-left">Short URL</th>
            <th className="border px-4 py-2 text-left">User ID</th>
            <th className="border px-4 py-2 text-left">Action</th>
          </tr>
        </thead>
        <tbody>
          {urls.map((item, index) => (
            <tr key={item.ID} className="hover:bg-gray-50">
              <td className="border px-4 py-2">{index + 1}</td>
              <td className="border px-4 py-2 truncate max-w-xs">
                <a href={item.Url} target="_blank" className="text-blue-600 hover:underline">
                  {item.Url}
                </a>
              </td>
              <td className="border px-4 py-2">
                <a
                  href={`http://13.213.2.17:8080/${item.ShortUrl}`}
                  target="_blank"
                  className="text-blue-600 hover:underline"
                >
                  {item.ShortUrl}
                </a>
              </td>
              <td className="border px-4 py-2">{item.UserID}</td>
              <td className="border px-4 py-2">
                <button
                  onClick={() => handleDelete(item.ShortUrl)}
                  className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
