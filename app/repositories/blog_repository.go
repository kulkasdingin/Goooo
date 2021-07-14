package repositories

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kulkasdingin/goooo/app/models"
)

var DB *gorm.DB

func FindBlogById(id int) (models.Blog, error) {
	var blog models.Blog

	var user models.User

	result := DB.First(&blog, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return blog, result.Error
	}

	DB.Model(&blog).Related(&user)

	safeUser := user.UserToSafeUser()

	blog.User = safeUser

	return blog, nil
}

func CreateBlog(blog models.Blog) models.Blog {
	DB.Create(&blog)

	return blog
}

func UpdateBlog(new_blog models.Blog, old_blog models.Blog) (models.Blog, error) {
	old_blog.Title = new_blog.Title
	old_blog.Content = new_blog.Content
	old_blog.HeaderImage = new_blog.HeaderImage

	DB.Save(&old_blog)

	new_blog, _ = FindBlogById(int(old_blog.ID))

	return new_blog, nil
}

func DeleteBlog(blog models.Blog) error {
	DB.Delete(&blog)

	return nil
}
