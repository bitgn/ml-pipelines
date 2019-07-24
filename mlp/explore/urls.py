from django.urls import path

from explore import views

app_name='explore'

urlpatterns = [
    path('datasets/', views.explore_datasets, name='explore_datasets'),
    path('datasets/<str:dataset_id>/', views.view_dataset, name='view_dataset'),
    path('datasets/<str:dataset_id>/lineage.svg', views.view_dataset_lineage, name='view_dataset_lineage'),

    path('', views.list_projects, name='projects'),
    path('projects/<str:project_id>/', views.view_project, name='view_project')

]