<template>
    <section
        id="contact"
        class="bg-primary-700 py-16"
    >
        <div class="max-w-2xl mx-auto text-center px-4 ">
            <h3 class="text-3xl font-extrabold text-white sm:text-4xl">
                {{ $t('contact.title') }}
            </h3>
            <p class="mt-4 text-lg leading-6 text-primary-200">
                {{ $t('contact.subtitle') }}
            </p>
        </div>

        <form
            class=" max-w-2xl mx-auto p-4 space-y-8"
            @submit.prevent="handleFormSubmit"
        >
            <label class="block">
                <span class="text-white text-lg">{{ $t('contact.email') }}</span>
                <input
                    v-model="email"
                    class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                    type="email"
                >
            </label>
            <label class="block">
                <span class="text-white text-lg">{{ $t('contact.message') }}</span>
                <textarea
                    v-model="message"
                    class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                />
            </label>
            <div class="w-full flex flex-col justify-center items-center">
                <button
                    class="bg-white px-4 py-2 mx-auto rounded-md shadow-sm text-primary-600 text-lg font-medium focus:outline-none focus:ring-2 focus:ring-white"
                    type="submit"
                >
                    {{ $t('contact.button') }}
                </button>

                <p
                    v-if="success"
                    class="mt-4 text-center text-lg leading-6 text-green-500"
                >
                    {{ $t('contact.success') }}
                </p>

                <p
                    v-if="error"
                    class="mt-4 text-lg leading-6 text-red-500"
                >
                    {{ error }}
                </p>
            </div>
        </form>
    </section>
</template>

<script setup>
const email = ref('')
const message = ref('')
const success = ref(false)
const error = ref(null)

const handleFormSubmit = async () => {
    error.value = null
    success.value = false
    try {
        const response = await fetch('https://formspree.io/f/xeoeylld', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: email.value,
                message: message.value,
            }),
        })

        if (!response.ok) {
            const res = await response.json()
            error.value = res.error || response.status
            return
        }

        email.value = ''
        message.value = ''
        success.value = true
    }
    catch (error) {
        error.value = error
    }
}
</script>
