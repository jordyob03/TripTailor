package DBmodels

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/lib/pq"
)

type Image struct {
	ImageId   int      `json:"imageId"`
	ImageData []byte   `json:"imageData"`
	Metadata  []string `json:"metadata"`
}

func InitializeImageTable(DB *sql.DB) error {
	profile := Image{
		ImageId:   1,
		ImageData: WebImageToByte("https://militaryhealthinstitute.org/wp-content/uploads/sites/37/2021/08/blank-profile-picture-png.png"),
		Metadata:  []string{"profile"},
	}

	cover := Image{
		ImageId:   2,
		ImageData: WebImageToByte("https://www.wallpapertip.com/wmimgs/1-19111_facebook-cover-photo-quotes-hd.jpg"),
		Metadata:  []string{"cover"},
	}

	id0, err := AddImage(DB, profile)
	if err != nil {
		return err
	} else {
		fmt.Printf("Profile at Id: %d added successfully!\n", id0)
	}

	id1, err := AddImage(DB, cover)
	if err != nil {
		return err
	} else {
		fmt.Printf("Cover at Id: %d added successfully!\n", id1)
	}

	return nil
}

func CreateImageTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS images (
		imageId SERIAL PRIMARY KEY,
		imageData BYTEA NOT NULL,
		metadata TEXT[]
	);`

	return CreateTable(DB, createTableSQL)
}

func AddImage(DB *sql.DB, image Image) (int, error) {
	insertImageSQL := `
	INSERT INTO images (imageData, metadata)
	VALUES ($1, $2)
	RETURNING imageId;`

	if image.Metadata == nil {
		image.Metadata = []string{}
	}

	var imageId int
	err := DB.QueryRow(
		insertImageSQL, image.ImageData, pq.Array(image.Metadata),
	).Scan(&imageId)

	return imageId, err
}

func GetImage(DB *sql.DB, imageId int) (Image, error) {
	getImageSQL := `
	SELECT * FROM images
	WHERE imageId = $1;`

	var image Image
	err := DB.QueryRow(getImageSQL, imageId).Scan(
		&image.ImageId, &image.ImageData, pq.Array(&image.Metadata),
	)

	if image.Metadata == nil {
		image.Metadata = []string{}
	}

	return image, err
}

func RemoveImage(DB *sql.DB, imageId int) error {
	removeImageSQL := `
	DELETE FROM images
	WHERE imageId = $1;`

	_, err := DB.Exec(removeImageSQL, imageId)

	return err
}

func AddImageMetaData(DB *sql.DB, imageId int, metadata string) error {
	return AddArrayAttribute(DB, "images", "imageId", imageId, "metadata", []string{metadata})
}

func RemoveImageMetaData(DB *sql.DB, imageId int, metadata string) error {
	return RemoveArrayAttribute(DB, "images", "imagesId", imageId, "metadata", metadata)
}

func ImageToByte(imagePath string) []byte {
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Printf("failed to read image: %s", err)
		return nil
	}
	return imageData
}

func WebImageToByte(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to download image: %s", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status: %s", resp.Status)
		return nil
	}

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read image data: %s", err)
		return nil
	}
	return imageData
}

func ByteToImage(imageData []byte, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(imageData)
	if err != nil {
		return err
	}

	fmt.Printf("Image saved to %s\n", outputPath)

	return nil
}

func ImageHandler(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/images/"):]
		imageID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid image ID", http.StatusBadRequest)
			return
		}

		image, err := GetImage(DB, imageID)
		if err != nil {
			http.Error(w, "Image not found", http.StatusNotFound)
			fmt.Printf("Error getting image: %s\n", err)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%d.png\"", imageID))
		w.WriteHeader(http.StatusOK)

		w.Write(image.ImageData)
	}
}
