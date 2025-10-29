from django.shortcuts import render, redirect
from django.views import View


# Create your views here.
class HomeView(View):
    def get(self, request):
        
        if request.headers.get("HX-Request"):
            # HTMX request - return just the content
            return render(request, 'home.html')
        
        # Direct page access - or full page refresh
        context = {
            'initial_content': 'home.html',
            'active_view': 'home'
        }
        return render(request, 'home.html', context)
    
class AchievementsView(View):
    def get(self, request):
        
        if request.headers.get("HX-Request"):
            # HTMX request - return just the content
            return render(request, 'achievements.html')
        
        # Direct page access - or full page refresh
        context = {
            'initial_content': 'achievements.html',
            'active_view': 'achievements'
        }
        return render(request, 'achievements.html', context)