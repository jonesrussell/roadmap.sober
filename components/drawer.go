package components

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
)

func Drawer() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := fmt.Fprintf(w, `
        <div class="relative z-10 hidden" id="slide-over" aria-labelledby="slide-over-title" role="dialog" aria-modal="true">
            <!-- Background backdrop, show/hide based on slide-over state. -->
            <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
            <div class="fixed inset-0 overflow-hidden">
                <div class="absolute inset-0 overflow-hidden">
                    <div class="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10">
                        <!-- Slide-over panel, show/hide based on slide-over state. -->
                        <div class="pointer-events-auto relative w-screen max-w-md">
                            <!-- Close button, show/hide based on slide-over state. -->
                            <div class="absolute left-0 top-0 -ml-8 flex pr-2 pt-4 sm:-ml-10 sm:pr-4">
                                <button id="close-panel" type="button" class="relative rounded-md text-gray-300 hover:text-white focus:outline-none focus:ring-2 focus:ring-white">
                                    <span class="absolute -inset-2.5"></span>
                                    <span class="sr-only">Close panel</span>
                                    <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"></path>
                                    </svg>
                                </button>
                            </div>
                            <div class="flex h-full flex-col overflow-y-scroll bg-white py-6 shadow-xl">
                                <div class="px-4 sm:px-6">
                                    <h2 class="text-base font-semibold leading-6 text-gray-900" id="slide-over-title">Panel title</h2>
                                </div>
                                <div class="relative mt-6 flex-1 px-4 sm:px-6" hx-swap="innerHTML" hx-target="#slide-over-content">
                                    <!-- Your content will be dynamically loaded here -->
                                    <div id="slide-over-content"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <script>
            const groupElements = document.querySelectorAll('[id^="group-"]');
            const closePanelButton = document.getElementById('close-panel');
            const slideOver = document.getElementById('slide-over');

            groupElements.forEach(group => {
                group.addEventListener('click', () => {
                    slideOver.classList.remove('hidden');
                });
            });

            closePanelButton.addEventListener('click', () => {
                slideOver.classList.add('hidden');
            });
        </script>
        `)
		return err
	})
}
