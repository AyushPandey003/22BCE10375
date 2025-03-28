export interface User {
    id: number;
    name: string;
    postCount: number;
  }
  
  export interface Post {
    id: number;
    user: string;
    content: string;
    commentCount: number;
    timestamp: number;
  }
  