// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Navigation() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"relative w-full px-8 text-gray-700 bg-white body-font\" data-tails-scripts=\"//unpkg.com/alpinejs\"><div class=\"container flex flex-col flex-wrap items-center justify-between py-5 mx-auto md:flex-row max-w-7xl\"><a href=\"#_\" class=\"relative z-10 flex items-center w-auto text-2xl font-extrabold leading-none text-black select-none\">Project</a><nav class=\"top-0 left-0 z-0 flex items-center justify-center w-full h-full py-5 -ml-0 space-x-5 text-base md:-ml-5 md:py-0 md:absolute\"><a href=\"#_\" class=\"relative font-medium leading-6 text-gray-600 transition duration-150 ease-out hover:text-gray-900\" x-data=\"{ hover: false }\" @mouseenter=\"hover = true\" @mouseleave=\"hover = false\"><span class=\"block\">Home</span> <span class=\"absolute bottom-0 left-0 inline-block w-full h-0.5 -mb-1 overflow-hidden\"><span x-show=\"hover\" class=\"absolute inset-0 inline-block w-full h-1 h-full transform bg-gray-900\" x-transition:enter=\"transition ease duration-200\" x-transition:enter-start=\"scale-0\" x-transition:enter-end=\"scale-100\" x-transition:leave=\"transition ease-out duration-300\" x-transition:leave-start=\"scale-100\" x-transition:leave-end=\"scale-0\"></span></span></a> <a href=\"#_\" class=\"relative font-medium leading-6 text-gray-600 transition duration-150 ease-out hover:text-gray-900\" x-data=\"{ hover: false }\" @mouseenter=\"hover = true\" @mouseleave=\"hover = false\"><span class=\"block\">Features</span> <span class=\"absolute bottom-0 left-0 inline-block w-full h-0.5 -mb-1 overflow-hidden\"><span x-show=\"hover\" class=\"absolute inset-0 inline-block w-full h-1 h-full transform bg-gray-900\" x-transition:enter=\"transition ease duration-200\" x-transition:enter-start=\"scale-0\" x-transition:enter-end=\"scale-100\" x-transition:leave=\"transition ease-out duration-300\" x-transition:leave-start=\"scale-100\" x-transition:leave-end=\"scale-0\"></span></span></a> <a href=\"#_\" class=\"relative font-medium leading-6 text-gray-600 transition duration-150 ease-out hover:text-gray-900\" x-data=\"{ hover: false }\" @mouseenter=\"hover = true\" @mouseleave=\"hover = false\"><span class=\"block\">Pricing</span> <span class=\"absolute bottom-0 left-0 inline-block w-full h-0.5 -mb-1 overflow-hidden\"><span x-show=\"hover\" class=\"absolute inset-0 inline-block w-full h-1 h-full transform bg-gray-900\" x-transition:enter=\"transition ease duration-200\" x-transition:enter-start=\"scale-0\" x-transition:enter-end=\"scale-100\" x-transition:leave=\"transition ease-out duration-300\" x-transition:leave-start=\"scale-100\" x-transition:leave-end=\"scale-0\"></span></span></a> <a href=\"#_\" class=\"relative font-medium leading-6 text-gray-600 transition duration-150 ease-out hover:text-gray-900\" x-data=\"{ hover: false }\" @mouseenter=\"hover = true\" @mouseleave=\"hover = false\"><span class=\"block\">Blog</span> <span class=\"absolute bottom-0 left-0 inline-block w-full h-0.5 -mb-1 overflow-hidden\"><span x-show=\"hover\" class=\"absolute inset-0 inline-block w-full h-1 h-full transform bg-gray-900\" x-transition:enter=\"transition ease duration-200\" x-transition:enter-start=\"scale-0\" x-transition:enter-end=\"scale-100\" x-transition:leave=\"transition ease-out duration-300\" x-transition:leave-start=\"scale-100\" x-transition:leave-end=\"scale-0\"></span></span></a></nav><div class=\"relative z-10 inline-flex items-center space-x-3 md:ml-5 lg:justify-end\"><a href=\"/login\" class=\"inline-flex items-center justify-center px-4 py-2 text-base font-medium leading-6 text-gray-600 whitespace-no-wrap bg-white border border-gray-200 rounded-md shadow-sm hover:bg-gray-50 focus:outline-none focus:shadow-none\" data-rounded=\"rounded-md\">Sign in</a> <span class=\"inline-flex rounded-md shadow-sm\"><a href=\"/signup\" class=\"inline-flex items-center justify-center px-4 py-2 text-base font-medium leading-6 text-white whitespace-no-wrap bg-blue-600 border border-blue-700 rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500\" data-rounded=\"rounded-md\" data-primary=\"blue-600\">Sign up</a></span></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate