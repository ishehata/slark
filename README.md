# Slark

Slark is an admin panel like django-admin built for Go.

Slark makes it easy to quickly manage your web app models and save time focusing
 on your product instead of consuming time to write a control panel.


    struct UserModel {
        ID string
        Name string
    }

    struct TagModel {
        ID string
        Title string
    }

    struct PostModel {
        ID string
        Title string
        Content string
        Author UserModel `relation:belongsTo`
        Tags []TagModel `relation:hasMany`
    }

    Slark.Register(UserModel)

    Slark.Register(PostModel)
