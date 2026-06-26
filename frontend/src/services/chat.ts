import { api } from './api'
import type {
  ChatMessage,
  ChatMessagePayload,
  ChatMessagesResponse,
  ChatRoomResponse,
  ChatRoomsResponse,
  MarkRoomReadPayload,
} from '../types/chat'

export async function openSchoolChatRoom() {
  const { data } = await api.post<ChatRoomResponse>('/chat/school/open')
  return data.room
}

export async function getChatRooms() {
  const { data } = await api.get<ChatRoomsResponse>('/chat/rooms')
  return data.rooms ?? []
}

export async function getMessages(
  roomId: string,
  params: { limit?: number; before?: string | null } = {},
) {
  const { data } = await api.get<ChatMessagesResponse>(`/chat/rooms/${roomId}/messages`, {
    params: {
      limit: params.limit,
      before: params.before || undefined,
    },
  })
  return {
    ...data,
    messages: data.messages ?? [],
  }
}

export async function sendMessage(roomId: string, content: string) {
  const payload: ChatMessagePayload = { content }
  const { data } = await api.post<ChatMessage>(`/chat/rooms/${roomId}/messages`, payload)
  return data
}

export async function markRoomRead(roomId: string, payload: MarkRoomReadPayload = {}) {
  const { data } = await api.patch<{ message: string }>(`/chat/rooms/${roomId}/read`, payload)
  return data
}
