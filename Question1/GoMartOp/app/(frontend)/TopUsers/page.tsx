"use client";
import { useEffect, useState } from "react";
import { User } from "@/types/interfaces"; // Import the User interface

export default function TopUsers() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/users")
      .then((res) => res.json())
      .then((data) => setUsers(data))
      .catch((err) => console.error("Error fetching users:", err));
  }, []);

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100">
      <h1 className="text-3xl font-bold text-gray-900 mt-6">Top Users</h1>
      <div className="mt-4 w-1/2 bg-white shadow-lg rounded-lg p-4">
        {users.map((user) => (
          <div key={user.id} className="border-b last:border-none p-4">
            <h2 className="font-semibold text-lg">{user.name}</h2>
            <p className="text-gray-700">Posts: {user.postCount}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
