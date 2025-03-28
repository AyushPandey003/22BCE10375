"use client";
import { useEffect, useState } from "react"; //use effects and use state are react hooks
import { Post } from "@/types/interfaces"; // Import Post interface

export default function TrendingPosts() {
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/posts?type=popular") //calling api from backend
      .then((res) => res.json())
      .then((data) => setPosts(data))
      .catch((err) => console.error("Error fetching trending posts:", err));
  }, []);

  return (
    <div className="flex flex-col items-center min-h-screen bg-yellow-100 p-6">
      <h1 className="text-3xl font-bold text-gray-900">Trending Posts</h1>
      <div className="mt-6 w-full max-w-3xl">
        {posts.length === 0 ? (
          <p className="text-gray-700 text-center">No trending posts found.</p>
        ) : (
          posts.map((post) => (
            <div key={post.id} className="bg-white shadow-lg rounded-lg p-4 mb-4">
              <h2 className="font-semibold text-xl">{post.user}</h2>
              <p className="text-gray-800 mt-2">{post.content}</p>
              <p className="text-gray-500 text-sm mt-2">Comments: {post.commentCount}</p>
            </div>
          ))
        )}
      </div>
    </div>
  );
}
