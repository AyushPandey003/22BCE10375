import React from "react";
import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="w-full bg-lime-600 p-4 shadow-md">
      <ul className="flex justify-center space-x-6">
        <li><Link href="/TopUsers" className="text-white font-semibold">Top Users</Link></li>
        <li><Link href="/TrendingPost" className="text-white font-semibold">Trending Posts</Link></li>
        <li><Link href="/Feed" className="text-white font-semibold">Feed</Link></li>   
      </ul>
    </nav>
  );
}
