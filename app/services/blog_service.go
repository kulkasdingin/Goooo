package services

import (
	"github.com/kulkasdingin/goooo/app/models"
	"github.com/kulkasdingin/goooo/app/repositories"
)

func GetBlogById(id int) (models.Blog, error) {
	return repositories.FindBlogById(id)

}

func CreateBlogFromRequest(blog models.Blog) models.Blog {
	return repositories.CreateBlog(blog)
}

func UpdateBlogFromRequest(new_blog models.Blog, old_id int) (models.Blog, error) { // TODO: Kudu check new_blog variablenya udah bener atau belum
	old_blog, err := GetBlogById(old_id)

	if err != nil {
		return new_blog, err
	}

	return repositories.UpdateBlog(new_blog, old_blog)
}

func DeleteBlogFromRequest(id int) error {
	blog, err := GetBlogById(id)

	if err != nil {
		return err
	}

	return repositories.DeleteBlog(blog)
}
