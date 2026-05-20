<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { PhArrowRight, PhGraduationCap } from "@phosphor-icons/vue";
import { dashboardByRole } from "../../router";
import { useAuthStore } from "../../stores/auth";

const auth = useAuthStore();
const route = useRoute();
const router = useRouter();

const email = ref("");
const password = ref("");
const isSubmitting = ref(false);
const errorMessage = ref("");

const canSubmit = computed(
  () => email.value.trim() !== "" && password.value.trim() !== "",
);

async function submit() {
  if (!canSubmit.value || isSubmitting.value) return;
  isSubmitting.value = true;
  errorMessage.value = "";

  try {
    await auth.login({ email: email.value, password: password.value });
    const role = auth.primaryRole();
    const fallback = role ? dashboardByRole[role] : "/unauthorized";
    await router.push((route.query.redirect as string | undefined) ?? fallback);
  } catch {
    errorMessage.value = "Email atau password tidak valid.";
  } finally {
    isSubmitting.value = false;
  }
}
</script>

<template>
  <section
    class="grid w-full max-w-5xl overflow-hidden rounded-[28px] bg-white shadow-[0_28px_80px_rgba(79,70,229,0.10)] md:grid-cols-[1.05fr_0.95fr]"
  >
    <div class="bg-[#eef2ff] px-8 py-10 sm:px-10">
      <div class="mb-12 flex items-center gap-3">
        <div
          class="flex h-11 w-11 items-center justify-center rounded-2xl bg-[#4f46e5] text-white"
        >
          <PhGraduationCap :size="24" weight="duotone" />
        </div>
        <div>
          <p class="text-sm font-medium text-[#4f46e5]">Eduverse</p>
          <p class="text-xs text-[#6b7280]">Academic workspace</p>
        </div>
      </div>

      <div class="max-w-md">
        <p class="mb-3 text-sm font-medium text-[#6b7280]">Multi-school LMS</p>
        <h1
          class="text-4xl font-medium leading-tight text-[#171322] sm:text-5xl"
        >
          Masuk ke ruang belajar yang lebih tenang.
        </h1>
        <p class="mt-5 text-sm leading-6 text-[#6b6475]">
          Satu login untuk siswa, guru, admin sekolah, dan super admin. Eduverse
          akan memilih ruang kerja berdasarkan role dan konteks sekolah.
        </p>
      </div>
    </div>

    <div class="px-8 py-10 sm:px-10">
      <div class="mb-8">
        <h2 class="text-2xl font-medium text-[#171322]">Login</h2>
        <p class="mt-2 text-sm text-[#7a7385]">
          Gunakan akun Eduverse yang sudah terdaftar.
        </p>
      </div>

      <form class="space-y-5" @submit.prevent="submit">
        <label class="block">
          <span class="mb-2 block text-sm text-[#5f5968]">Email</span>
          <input
            v-model="email"
            class="h-12 w-full rounded-2xl border border-[#e7e2da] bg-[#fbfaf8] px-4 text-sm outline-none transition focus:border-[#4f46e5] focus:bg-white"
            type="email"
            autocomplete="email"
            placeholder="nama@sekolah.sch.id"
          />
        </label>

        <label class="block">
          <span class="mb-2 block text-sm text-[#5f5968]">Password</span>
          <input
            v-model="password"
            class="h-12 w-full rounded-2xl border border-[#e7e2da] bg-[#fbfaf8] px-4 text-sm outline-none transition focus:border-[#4f46e5] focus:bg-white"
            type="password"
            autocomplete="current-password"
            placeholder="••••••••"
          />
        </label>

        <p
          v-if="errorMessage"
          class="rounded-2xl bg-[#fff1f0] px-4 py-3 text-sm text-[#b42318]"
        >
          {{ errorMessage }}
        </p>

        <button
          class="flex h-12 w-full items-center justify-center gap-2 rounded-2xl bg-[#4f46e5] text-sm font-medium text-white transition hover:bg-[#4338ca] disabled:cursor-not-allowed disabled:bg-[#bab7d8]"
          type="submit"
          :disabled="!canSubmit || isSubmitting"
        >
          {{ isSubmitting ? "Memproses..." : "Masuk" }}
          <PhArrowRight v-if="!isSubmitting" :size="18" />
        </button>
      </form>
    </div>
  </section>
</template>
