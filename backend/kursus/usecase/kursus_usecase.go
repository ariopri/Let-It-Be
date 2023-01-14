package usecase

import (
	"context"

	"time"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type kursus struct {
	kursusRepo models.KursusRepository
	userRepo   models.UserRepository
	ctxTimeout time.Duration
}

func NewKursusUseCase(kursusRepo models.KursusRepository, userRepo models.UserRepository, ctxTimeout time.Duration) models.KursusUseCase {
	return &kursus{
		kursusRepo: kursusRepo,
		userRepo:   userRepo,
		ctxTimeout: ctxTimeout,
	}
}

func (k *kursus) getUsers(ctx context.Context, result []models.KursusResp) ([]models.KursusResp, error) {
	g, c := errgroup.WithContext(ctx)

	mapUser := map[int64]models.User{}

	for _, kursus := range result {
		mapUser[int64(kursus.User.ID)] = models.User{}
	}

	chanUser := make(chan models.User)
	for userID := range mapUser {
		userID := userID
		g.Go(func() error {
			res, err := k.userRepo.GetById(c, uint(userID))
			if err != nil {
				return err
			}
			chanUser <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanUser)
	}()

	for user := range chanUser {
		if user != (models.User{}) {
			mapUser[int64(user.ID)] = user
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for index, item := range result {
		if u, ok := mapUser[int64(item.User.ID)]; ok {
			result[index].User = u
		}
	}

	return result, nil
}

func (k *kursus) GetAll(ctx context.Context) ([]models.KursusResp, error) {
	c, cancel := context.WithTimeout(ctx, k.ctxTimeout)

	defer cancel()

	res, err := k.kursusRepo.GetAll(c)
	if err != nil {
		return []models.KursusResp{}, err
	}

	res, err = k.getUsers(c, res)
	if err != nil {
		return []models.KursusResp{}, err
	}

	return res, nil
}

func (k *kursus) GetById(ctx context.Context, id int64) (models.KursusResp, error) {
	c, cancel := context.WithTimeout(ctx, k.ctxTimeout)

	defer cancel()

	res, err := k.kursusRepo.GetById(c, id)
	if err != nil {
		return models.KursusResp{}, err
	}

	user, err := k.userRepo.GetById(c, res.User.ID)
	if err != nil {
		return models.KursusResp{}, err
	}

	res.User = user

	return res, nil
}

func (k *kursus) Store(ctx context.Context, kursus *models.Kursus) (models.KursusResp, error) {
	c, cancel := context.WithTimeout(ctx, k.ctxTimeout)

	defer cancel()

	res, err := k.kursusRepo.Store(c, kursus)
	if err != nil {
		return models.KursusResp{}, err
	}

	user, err := k.userRepo.GetById(c, res.User.ID)
	if err != nil {
		return models.KursusResp{}, err
	}

	res.User = user

	return res, nil
}

func (k *kursus) Update(ctx context.Context, id int64, kursus *models.Kursus) (models.KursusResp, error) {
	c, cancel := context.WithTimeout(ctx, k.ctxTimeout)

	defer cancel()

	res, err := k.kursusRepo.Update(c, id, kursus)
	if err != nil {
		return models.KursusResp{}, err
	}

	user, err := k.userRepo.GetById(c, res.User.ID)
	if err != nil {
		return models.KursusResp{}, err
	}

	res.User = user

	return res, nil
}

func (k *kursus) Delete(ctx context.Context, id int64) error {
	c, cancel := context.WithTimeout(ctx, k.ctxTimeout)

	defer cancel()

	err := k.kursusRepo.Delete(c, id)
	if err != nil {
		return err
	}

	return nil
}
