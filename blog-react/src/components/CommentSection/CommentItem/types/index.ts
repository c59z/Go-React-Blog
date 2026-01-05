import type { Comment } from "@/api/comment";

export interface CommentItemProps {
  comment: Comment;
  level: 1 | 2 | 3;
  canDelete: boolean;

  /** 事件回调 */
  onReply: (commentId: string) => void;
  onLike: (commentId: string) => void;
  onDelete: (commentId: string) => void;
}
