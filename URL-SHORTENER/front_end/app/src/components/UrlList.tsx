import { useEffect, useState } from "react";
import React from 'react'
import axios from "axios";

interface UrlItem {
  ID: number;
  Url: string;
  ShortUrl: string;
  UserID: string;
}
type Props = {
  clicked: boolean,
  children?: React.ReactNode
}
export default function UrlList({ clicked }: Props) {
  const [urls, setUrls] = useState<UrlItem[]>([]);

  useEffect(() => {
    const fetchUrls = async () => {
      try {
        const res = await axios.get<UrlItem[]>("http://13.213.2.17:80/urls");
        setUrls(res.data);
      } catch (err) {
        console.error("Failed to fetch URLs:", err);
      }
    };
    fetchUrls();
  }, [clicked]);

  return (
    <div className="mt-8 bg-white p-6 rounded-2xl shadow-lg w-full max-w-3xl mx-auto">
      <h2 className="text-xl font-semibold text-gray-700 mb-4">URL List</h2>

      <table className="w-full border-collapse border border-gray-300 text-sm">
        <thead className="bg-gray-100">
          <tr>
            <th className="border border-gray-300 px-4 py-2 text-left">#</th>
            <th className="border border-gray-300 px-4 py-2 text-left">Original URL</th>
            <th className="border border-gray-300 px-4 py-2 text-left">Short URL</th>
            <th className="border border-gray-300 px-4 py-2 text-left">User ID</th>
          </tr>
        </thead>
        <tbody>
          {urls.map((item, index) => (
            <tr key={item.ID} className="hover:bg-gray-50">
              <td className="border border-gray-300 px-4 py-2">{index + 1}</td>
              <td className="border border-gray-300 px-4 py-2 truncate max-w-xs">
                <a href={item.Url} target="_blank" rel="noopener noreferrer" className="text-blue-600 hover:underline">
                  {item.Url}
                </a>
              </td>
              <td className="border border-gray-300 px-4 py-2">
                <a
                  href={`http://13.213.2.17/${item.ShortUrl}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-blue-600 hover:underline"
                >
                  {item.ShortUrl}
                </a>
              </td>
              <td className="border border-gray-300 px-4 py-2">{item.UserID}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
