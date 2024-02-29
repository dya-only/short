<script lang="ts">
import { onMount } from 'svelte'

export let key: string

onMount(async () => {
  const res = await fetch(`/api?key=${key}`, { method: 'GET' })
  const data = await res.json()

  window.location.href = data.url
})
</script>