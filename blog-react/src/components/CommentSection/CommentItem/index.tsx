import { useState } from "react";
import type { CommentItemProps } from "./types";
import "./index.scss";
import { commentCreate } from "@/api/comment";

const CommentItem = ({
  comment,
  level,
  onDelete,
  canDelete,
  onReplySuccess,
}: CommentItemProps) => {
  const { id, user, content, created_at } = comment;

  const [isReplying, setIsReplying] = useState(false);
  const [replyContent, setReplyContent] = useState("");

  const handleReply = () => {
    setIsReplying(true);
  };

  const handleSend = async () => {
    if (!replyContent.trim()) return;

    try {
      const res = await commentCreate({
        content: replyContent,
        article_id: comment.article_id,
        p_id: comment.id,
      });

      if (res.code === 0) {
        onReplySuccess?.();
        setReplyContent("");
        setIsReplying(false);
      }
    } catch (err) {
      console.error("回复失败", err);
    }
  };

  return (
    <div className={`article-comment-item level-${level}`}>
      <img
        className="comment-avatar"
        src={`${import.meta.env.VITE_SERVER_URL}${user.avatar}`}
        alt={user.username}
      />

      <div className="comment-body">
        <div className="comment-header">
          <span className="comment-username">{user.username}</span>
          <span className="comment-time">
            {new Date(created_at).toLocaleString()}
          </span>

          {canDelete && (
            <button
              className="comment-delete"
              onClick={() => onDelete(id.toString())}
            >
              删除
            </button>
          )}
        </div>

        <div className="comment-content">{content}</div>

        <div className="comment-actions">
          <button onClick={handleReply}>回复</button>
        </div>

        {isReplying && (
          <div className="comment-reply-box">
            <textarea
              value={replyContent}
              onChange={(e) => setReplyContent(e.target.value)}
              placeholder="输入你的回复..."
            />
            <div className="reply-buttons">
              <button className="send" onClick={handleSend}>
                发送
              </button>
              <button className="cancel" onClick={() => setIsReplying(false)}>
                取消
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default CommentItem;
