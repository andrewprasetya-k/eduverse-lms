export type CommentSourceType = "feed" | "material" | "assignment";

export interface CommentItem {
  commentId: string;
  sourceType: CommentSourceType | string;
  sourceId: string;
  content: string;
  creatorName?: string;
  createdAt: string;
  updatedAt?: string;
  isMine?: boolean;
}

export interface GetCommentsParams {
  sourceType: CommentSourceType;
  sourceId: string;
}

export interface CreateCommentPayload {
  sourceType: CommentSourceType;
  sourceId: string;
  content: string;
}
