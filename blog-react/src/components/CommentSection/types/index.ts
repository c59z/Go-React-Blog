import type { Comment } from "@/api/comment";

export interface RenderComment {
  comment: Comment;
  level: 1 | 2 | 3;
}
