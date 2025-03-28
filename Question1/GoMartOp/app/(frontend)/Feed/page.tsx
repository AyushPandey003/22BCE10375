"use client";
import React, { useEffect, useState } from "react";
import { Post } from "@/types/interfaces"; // Import Post interface

export default function Feed() {

  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => { // use effect for client side rendering
    const fetchPosts = async () => {
      try {
        const response = await fetch("http://localhost:3000/api/livefeed");
        const data = await response.json();
        setPosts(data);
      } catch (error) {
        console.error("Error fetching feed:", error);
      }
    };

    fetchPosts();
    const interval = setInterval(fetchPosts, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100">
      <h1 className="text-3xl font-bold text-gray-900 mt-6">Live Feed</h1>
      <div className="mt-4 w-1/2 bg-white shadow-lg rounded-lg p-4">
        {posts.map((post) => (
          <div key={post.id} className="border-b last:border-none p-4">
            <h2 className="font-semibold text-lg">{post.user}</h2>
            <p className="text-gray-700">{post.content}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
