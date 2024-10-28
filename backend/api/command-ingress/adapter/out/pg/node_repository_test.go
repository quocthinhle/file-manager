package pgoutadapter_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	pgoutadapter "github.com/quocthinhle/file-manager-api/command-ingress/adapter/out/pg"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/entity"
)

var _ = Describe("NodeRepository", func() {
	var db *pgxpool.Pool
	var err error
	var nodeRepository *pgoutadapter.NodeOutputAdapter
	var ctx context.Context
	var cancelFn context.CancelFunc

	BeforeEach(func() {
		db, err = pgxpool.New(context.Background(), dbURL)
		nodeRepository = pgoutadapter.NewNodeOutputAdapter(db)
		Expect(err).NotTo(HaveOccurred())
		ctx, cancelFn = context.WithCancel(context.Background())
	})

	AfterEach(func() {
		db.Close()
		cancelFn()
	})

	Describe("create a directory", func() {
		It("should create a directory", func() {
			content := entity.Content{
				ID:          uuid.MustParse(gofakeit.UUID()),
				Name:        "root-1",
				Description: "Root dir folder 1",
				Type:        "directory",
				ParentID:    uuid.Nil,
				OwnerID:     uuid.MustParse(gofakeit.UUID()),
				Children:    nil,
			}

			contentCreated, err := nodeRepository.Create(ctx, content)
			Expect(err).NotTo(HaveOccurred())

			// Expect created node
			Expect(contentCreated).NotTo(BeNil())
			Expect(contentCreated.ID).To(Equal(content.ID))
			Expect(contentCreated.Name).To(Equal(content.Name))
			Expect(contentCreated.Type).To(Equal(content.Type))
			Expect(contentCreated.ParentID).To(Equal(content.ParentID))
			Expect(contentCreated.OwnerID).To(Equal(content.OwnerID))
			assertNodeClosure(ctx, db, content.ID, content.ID, 0)

			// Insert dir with parentID
			childContent := entity.Content{
				ID:          uuid.MustParse(gofakeit.UUID()),
				Name:        "child-root-1",
				Description: "child dir root 1",
				Type:        "directory",
				ParentID:    content.ID,
				OwnerID:     uuid.MustParse(gofakeit.UUID()),
				Children:    nil,
			}

			childContentCreated, err := nodeRepository.Create(ctx, childContent)
			Expect(err).NotTo(HaveOccurred())
			Expect(childContentCreated).NotTo(BeNil())
			Expect(childContentCreated.ID).To(Equal(childContent.ID))
			Expect(childContentCreated.Name).To(Equal(childContent.Name))
			Expect(childContentCreated.Type).To(Equal(childContent.Type))
			Expect(childContentCreated.ParentID).To(Equal(childContent.ParentID))
			Expect(childContentCreated.OwnerID).To(Equal(childContent.OwnerID))
			assertNodeClosure(ctx, db, childContent.ID, childContent.ID, 0)
			assertNodeClosure(ctx, db, content.ID, childContent.ID, 1)
		})
	})
})

func assertNodeClosure(
	ctx context.Context,
	conn *pgxpool.Pool,
	parentId uuid.UUID,
	childId uuid.UUID,
	depth int,
) {
	GinkgoHelper()

	var realParentId uuid.UUID
	var realAncestorId uuid.UUID
	var realDepth int

	err := conn.QueryRow(
		ctx, "SELECT ancestor_id, descendant_id, depth FROM node_closure WHERE ancestor_id = $1 AND descendant_id = $2 AND depth = $3",
		parentId,
		childId,
		depth,
	).Scan(&realParentId, &realAncestorId, &realDepth)

	Expect(err).NotTo(HaveOccurred())
	Expect(realParentId).To(Equal(parentId))
	Expect(realAncestorId).To(Equal(childId))
	Expect(realDepth).To(Equal(realDepth))
}
