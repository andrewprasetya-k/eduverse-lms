export interface ChatLastMessage {
  messageId: string
  senderId: string
  senderName: string
  content: string
  createdAt: string
}

export interface ChatRoom {
  roomId: string
  roomName: string
  roomType: string
  roomRefType: string
  roomRefId: string
  schoolId: string
  schoolName?: string
  lastMessage?: ChatLastMessage | null
  lastMessageAt?: string | null
  unreadCount: number
  canSend: boolean
}

export interface ChatMessage {
  messageId: string
  roomId: string
  senderId: string
  senderName: string
  senderRole: string
  content: string
  messageType: 'text' | string
  createdAt: string
  isMine: boolean
}

export interface ChatRoomsResponse {
  rooms: ChatRoom[]
}

export interface ChatRoomResponse {
  room: ChatRoom
}

export interface ChatMessagesResponse {
  messages: ChatMessage[]
  nextBefore?: string | null
  hasMore: boolean
}

export interface ChatMessagePayload {
  content: string
}

export interface MarkRoomReadPayload {
  lastReadMessageId?: string
}

