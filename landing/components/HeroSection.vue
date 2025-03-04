<template>
    <section class="relative bg-white overflow-hidden">
        <div class="max-w-7xl mx-auto">
            <div class="relative z-10 pb-8 bg-white sm:pb-16 md:pb-20 lg:max-w-2xl lg:w-full lg:pb-28 xl:pb-32">
                <svg
                    class="hidden lg:block absolute right-0 inset-y-0 h-full w-32 text-white transform translate-x-1/2"
                    fill="currentColor"
                    viewBox="0 0 100 100"
                    preserveAspectRatio="none"
                    aria-hidden="true"
                >
                    <polygon points="50,0 100,0 50,100 0,100" />
                </svg>

                <div class="pt-10 mx-auto max-w-7xl px-4 sm:pt-12 sm:px-6 md:pt-16 lg:pt-20 lg:px-8 xl:pt-28">
                    <div class="sm:text-center lg:text-left">
                        <h1
                            class="text-4xl tracking-tight font-extrabold text-gray-900 md:text-5xl drop-shadow"
                            v-html=" $t('hero.title')"
                        />
                        <p class="mt-3 text-base text-gray-500 sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0">
                            {{ $t('hero.subtitle') }}
                        </p>
                        <div class="mt-5 sm:mt-8 sm:flex sm:justify-center lg:justify-start">
                            <div class="rounded-md shadow-md">
                                <a
                                    href="#pricing"
                                    class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 md:py-4 md:text-lg md:px-10 transition-all"
                                >
                                    {{ $t('hero.button') }}
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="lg:absolute lg:inset-y-0 lg:right-0 lg:w-1/2">
                <img
                    class="h-96 mx-auto object-cover lg:h-full animate-fly"
                    src="@/assets/hero.jpg"
                    alt="Projectly banner"
                >
            </div>
        </div>
    </section>
</template>

<style>
@keyframes fly {
    0%, 100% {
        transform: translateY(-3%);
        animation-timing-function: cubic-bezier(0.8,0,1,1);
    }
    50% {
        transform: none;
        animation-timing-function: cubic-bezier(0,0,0.2,1);
    }
}
.animate-fly {
    animation: fly 5s infinite;
}
</style>
