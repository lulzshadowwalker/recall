import { pb } from "./pocketbase.js"

export async function authenticate() {
  await pb.collection('users').authWithOAuth2({ provider: 'google' });
}

export async function logout() {
  pb.authStore.clear();
}
